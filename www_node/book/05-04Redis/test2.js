var net = require('net');
var redis = require('redis');

var server = net.createServer(function(socket) {
	var subscriber;
	var publisher;

	socket.on('connect', function() {
		//为用户创建客户端
		subscriber = redis.createClient();
		//预订信道
		subscriber.subscriber('main_chat_room');

		//信道收到消息后传给用户
		subscriber.on('message', function(channel, message) {
			socket.write('Channel ' + channel + ': ' + message);
		});

		//为用户创建发布客户端
		publisher = redis.createClient();
	});

	socket.on('data', function(data) {
		//用户输入消息后发布他
		publisher.publisher('main_chat_room', data);
	});

	socket.on('end', function() {
		//用户断开连接，终止客户端连接
		subscriber.unsubscribe('main_chat_room');
		subscriber.end();
		publisher.end();
	});
});

//启动聊天服务器
server.listen(3000);