var md5 = require('js-md5');

//生成随机数
exports.getSalt = function(n, fn) {
	fn(false, Math.ceil(Math.random()*n));
}

//生成用户密码hash值
exports.hash = function(pass, n, fn) {
	var p = md5(md5(pass) + n);
	fn(false, p);
}
