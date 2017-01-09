var socketio = require('socket.io');
var io;
var guestNumber = 1;
var nickNames = {};
var namesUsed = [];
var currentRomm = {};

exports.listen = function(server) {
    io = socketio.listen(server);
    io.set('log level', 1);
    //定义每个用户连接的处理逻辑
    io.sockets.on('connection', function(socket) {
        //在用户连接上来时赋予一个访客名称
        guestNumber = assignGuestName(socket, guestNumber, nickNames, namesUsed);
        //用户连接上来时把他放入聊天室Lobby
        joinRoom(socket, 'Lobby');

        //处理用户的消息，更名，以及聊天室的创建和变更
        handleMessageBroadcasting(socket, nickNames);
        handleNameChangeAttempts(socket, nickNames, namesUsed);
        handleRoomJoining(socket);

        //用户发出请求时,向其提供已经被占用的聊天室列表
        socket.on('rooms', function() { //on注册事件
            socket.emit('rooms', io.sockets.manager.rooms); //触发事件
        })

        //定义用户断开连接后的清除逻辑
        handleClientDisconnection(socket, nickNames, namesUsed);
    })
}

//用户分配昵称
function assignGuestName(socket, guestNumber, nickNames, namesUsed) {
    var name = 'Guest' + guestNumber;
    nickNames[socket.id] = name; //把用户昵称与客户连接ID相关联
    //让用户知道他们的昵称
    socket.emit('nameResult', {
        success: true,
        name: name
    });
    namesUsed.push(name);
    return guestNumber + 1;
}

//将用户放入聊天室
function joinRoom(socket, room) {
	console.log(room);
	console.log(socket.id);
    //让用户进入房间
    socket.join(room);
    currentRomm[socket.id] = room;
    //让用户知道他们进入了新房间
    socket.emit('joinResult', {
        room: room
    });
    //让房间里的其他用户知道有新用户进入房间
    socket.broadcast.to(room).emit('message', {
        text: nickNames[socket.id] + ' has joined ' + room + '.'
    })

    var usersInRoom = io.sockets.clients(room);
    //获取房间中用户列表
    if (usersInRoom.length > 1) {
        var usersInRoomSummary = 'User currently in ' + room + ': ';
        for (var index in usersInRoom) {
            var userSocketId = usersInRoom[index].id;
            if (userSocketId != socket.id) {
                if (index > 0) {
                    usersInRoomSummary += ', ';
                }
                usersInRoomSummary += nickNames[userSocketId];
            }
        }
        usersInRoomSummary += '.';
        //将列表发送给该用户
        socket.emit('message', {
            text: usersInRoomSummary
        })
    }
}

function handleNameChangeAttempts(socket, nickNames, namesUsed) {

    socket.on('nameAttempt', function(name) {
        if (name.indexOf('Guest') == 0) { //用户名头部为Guest时不能更改
            socket.emit('nameResult', {
                success: false,
                message: 'Names cannot begin with "Guest".'
            });
        } else {
            if (namesUsed.indexOf(name) == -1) { //判断用户名是否重复
                var previousName = nickNames[socket.id];
                var previousNameIndex = namesUsed.indexOf(previousName);
                namesUsed.push(name);
                nickNames[socket.id] = name;
                delete namesUsed[previousNameIndex];

                socket.emit('nameResult', {
                    success: true,
                    name: name
                });
                //向该用户房间中用户发送更名信息
                socket.broadcast.to(currentRomm[socket.id]).emit('message', {
                    text: previousName + ' is now known as ' + name + '.'
                });
            } else {
                //发出已被占用的错误消息
                socket.emit('nameResult', {
                    success: false,
                    message: 'That name is already in use.'
                });
            }
        }
    });
}


//向房间中其他用户发送信息
function handleMessageBroadcasting(socket) {
    socket.on('message', function(message) {
        socket.broadcast.to(message.room).emit('message', {
            text: nickNames[socket.id] + ': ' + message.text
        })
    });
}

//创建房间
function handleRoomJoining(socket) {
    socket.on('join', function(room) {
        socket.level(currentRomm[socket.id]);
        joinRoom(socket, room.newRoom);
    });
}

//用户断开连接
function handleClientDisconnection(socket) {
    socket.on('disconnect', function() {
        var nameIndex = namesUsed.indexOf(nickNames[socket.id]);
        delete namesUsed[nameIndex];
        delete nickNames[socket.id];
    });
}
