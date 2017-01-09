//防止过大的恶意请求
var connect = require('connect');
var http = require('http');
var bodyParser = require('body-parser');
var requestLimit = require('connect-limit');

var app = connect()
			.use(requestLimit('32kb'))//请求不能超过32k
			.use(type('application/x-www-form-urlencoded', requestLimit('64kb')))
			.use(bodyParser());

http.createServer(app).listen(3000);