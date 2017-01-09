//http验证
var connect = require('connect');
var app = connect();

app.use(logger)
	//.use('/blog', blog)
	//.use('/posts', blog)
	.use('/admin', restrict)
	.use('/admin', admin)
	.use(hello);

app.listen(3000);

function logger(req, res, next) {
	console.log('%s %s', req.method, req.url);
	next();//将控制权交给下一个组件
}

function hello(req, res, next) {
	res.setHeader('Content-Type', 'text/plain');
	res.end('hello world');
}

function restrict(req, res, next) {
	//base64认证
	var authorizaztion = req.headers['authorization'];
	//if(!authorizaztion) return next(new Error('Unauthorized'));

	if(!authorizaztion) {
		res.statusCode = 401;
		res.setHeader('WWW-Authenticate', 'Basic realm="Secure Area"');
		res.end();
		return ;
	}


	var parts = authorizaztion.split(' ');
	var scheme = parts[0];
	var auth = new Buffer(parts[1], 'base64').toString().split(':');
	var user = auth[0];
	var pass = auth[1];

	//检查用户名密码
	authenticateWithDatabase(user, pass, function(err) {
		if(err) return next(err);//报告分析器出错
		next();
	});
}

function authenticateWithDatabase(user, pass, cb) {
	cb(false);
}

function admin(req, res, next) {
	switch(req.url) {
		case '/':
			res.end('try /users');
			break;
		case '/users':
			res.setHeader('Content-Type', 'application/json');
			res.end(JSON.stringify(['tobi', 'loki', 'jane']));
			break;
	}
}