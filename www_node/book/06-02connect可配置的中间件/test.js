var connect = require('connect');
//路由器组件
var router = require('./middleware/router');
//路由对象
var routes = {
	GET: {
		'/users': function(req, res) {
			res.end('tobi, loki, ferret');
		},
		'/user/:id': function(req, res, id) {
			res.end('user ' + id);
		}
	},
	DELETE: {
		'/user/:id': function(req, res, id) {
			res.end('delete user ' + id);
		}
	}
}

connect.use(router(routes)).listen(3000);