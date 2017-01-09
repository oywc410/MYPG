var connect = require('connect');
var cookieParser = require('cookie-parser');

var app = connect()
			.use(cookieParser())
			.use(function(req, res) {
				console.log(req.cookies);//cookies值  
				//curl http://localhost:3000/ -H "Cookie: foo=bar,bar=baz"
				console.log(req.signedCookies);//cookie签名

				res.end('hello\n');
			}).listen(3000);




var app = connect()
			.use(function(req, res) {
				res.setHeader('Set-Cookie', 'foo=bar');
				res.setHeader('Set-Cookie', 'tobi=ferret;Exprires=Tue, 08 Jun 2021 10:18:14 GMT');
			}).listen(3001);