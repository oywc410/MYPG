var connect = require('connect');
var app = connect();

app.use(logger).use(hello);//按顺序调用
//app.use(logger).use(用户认证).use(用户认证).use(hello);
//app.use(logger).use('/admin', restrict).use('/admin', admin).use(hello);
app.listen(3000);

function logger(req, res, next) {
	console.log('%s %s', req.method, req.url);
	next();//将控制权交给下一个组件
}

function hello(req, res, next) {
	res.setHeader('Content-Type', 'text/plain');
	res.end('hello world');
}