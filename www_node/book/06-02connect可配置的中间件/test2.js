var connect = require('connect');
//路由器组件
var router = require('./middleware/router');

connect()
	.use(router(require('./routes/user')))
	.use(router(require('./routes/admin')))
	.listen(3000);