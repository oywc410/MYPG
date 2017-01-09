var connect = require('connect');
var cookieParser = require('cookie-parser');
var session = require('connect-session');



var app = connect()
			.use(cookieParser('keyboard cat'))
			.use(session.session.Session())
			.use(function(req, res, next) {
				var sess = req.session;
				if(sess.views) {
					res.setHeader('Content-Type', 'text/html');
					res.write('<p>views: ' + sess.views + '</p>');
					res.end();
				} else {
					sess.views = 1;
					res.end('we');
				}
			});
app.listen(3000);