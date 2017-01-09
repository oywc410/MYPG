//伪造http method (GET,POST => 自定义)
//日志
var connect = require('connect');
var bodyParser = require('body-parser');
var methodOverride = require('method-override');

function edit(req, res, next) {
	if('GET' != req.method) return next();
	res.setHeader('Content-Type', 'text/html');
	res.write('<form method="post">');
	res.write('<input type="hidden" name="_method" value="put" />');
	res.write('<input type="text" name"user[name]" value="Tobi" />');
	res.write('<input type="submit" value="Update" />');
	res.write('</form>');
	res.end();
}

function update(req, res, next) {
	if('PUT' != req.method) return next();
	res.end('Updated name to ' + req.body.user.name);
}

var app = connect()
			.use(methodOverride())
			.use(bodyParser())
			.use(edit)
			.use(function(req, res, next) {
				console.log(req.method);
				console.log(req.originalMethod);
				next();
			})
			.use(update);
app.listen(3000);