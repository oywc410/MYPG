
var connect = require('connect');
var query = require('connect-query');
//GET请求解析

var app = connect()
			.use(query())
			.use(function(req, res, next) {
				res.setHeader('Content-Type', 'application/json');
				res.end(JSON.stringify(req.query));
			}).listen(3000);