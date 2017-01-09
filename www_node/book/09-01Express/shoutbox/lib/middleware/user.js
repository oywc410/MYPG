var User = require('../user');

//判断用户是否登入的中间件
module.exports = function(req, res, next) {

	if(req.remoteUser) {
		res.locals.user = req.remoteUser;
	}

	//从回话中取出已登录的用户ID
	var uid = req.session.uid;
	if(!uid) return next();
	//从数据库中查询信息
	User.get(uid, function(err, user) {
		if(err) return next();
		req.user = res.locals.user = user;
		next();
	});
}