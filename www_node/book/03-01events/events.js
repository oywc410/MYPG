var events = require('events');
var net = require('net');

var channel = new events.EventEmitter();
channel.clients = {};
channel.subscriptions = {};

channel.setMaxListeners(50);

channel.on('join', function(id, client) {

    //获取在线数
    var welcome = "Welcome!\n" + 'Guests online: ' + this.listeners('broadcast').length;
    client.write(welcome + "\n");

    //储存client对象
    this.clients[id] = client;
    //发送广播
    this.subscriptions[id] = function(senderId, message) {
            if (id != senderId) {
                this.clients[id].write(message);
            }
        }
        //监听事件
    this.on('broadcast', this.subscriptions[id]);
});

//停止服务
channel.on('shutdown', function() {
    channel.emit('broadcast', '', "Chat has shut down.\n");
    channel.removeAllListeners('broadcast');
});

//连接断开
channel.on('leave', function(id) {
    channel.process.removeListener('broadcast', this.subscriptions[id]);
    channel.emit('broadcast', id, id + " has left the chat.\n");
});

var server = net.createServer(function(client) {
    var id = client.remoteAddress + ':' + client.remoteProt;
    //当有用户连接到服务器上时发送一个join时间
    client.on('connect', function() {
        channel.emit('join', id, client);
    });

    //当有用户发送数据时,发送一个broadcast事件
    client.on('data', function(data) {
        data = data.toString();
        if (data == "shutdown\r\n") { //输入全部停止命令
            channel.emit('shutdown');
        }
        channel.emit('broadcast', id, data);
    });

    //用户断开连线时触发协议
    client.on('clost', function() {
        channel.emit('leave', id);
    });
});

server.listen(8888);
//telnet 127.0.0.1 8888
