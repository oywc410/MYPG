//解析请求主体POST
//将用户传来的值解析为json
var connect = require('connect');
var bodyParser = require('body-parser');

var app = connect()
			.use(bodyParser())
			.use(function(req, res) {
				console.log(req.body);
				console.log(req.files);
				res.end('thanks!');
			}).listen(3000);
//curl -d '{"username":"tobi"}' -H "Content-Type: application/json" http://localhost:3000/