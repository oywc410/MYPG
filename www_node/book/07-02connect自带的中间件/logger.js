//日志
var connect = require('connect');
var logger = require('connect-logger');

var app = connect()
			.use(logger())//默认
			//.use(logger('short'))//预定义log格式
			.use(logger({date: "YY.MM.DD HH:mm:ss", format: "%date %status %method %url (%route - %time)"}))//自定义
			.use(function(req, res, next) {
				res.end('logs');
			}).listen(3000);