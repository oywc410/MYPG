var redis = require('redis');
//var bcrypt = require('bcrypt');
var bcrypt = require('./bcrypt.js');

//创建到Redis的长链接
var db = redis.createClient();

module.exports = User;

function User(obj) {
    //遍历传入对象中的健
    for (var key in obj) {
        this[key] = obj[key];
    }
}

User.prototype.save = function(fn) {
    if (this.id) { //用户已经存在
        this.update(fn);
    } else {
        var user = this;
        db.incr('user:ids', function(err, id) { //创建唯一ID
            if (err) return fn(err);
            user.id = id; //设定ID
            user.hashPassword(function(err) {
                if (err) return fn(err);
                user.update(fn); //保持对象
            });
        });
    }
}

User.prototype.update = function(fn) {
    var user = this;
    var id = user.id;
    db.set('user:id:' + user.name, id, function(err) { //用名称索引用户ID
        if (err) return fn(err);
        db.hmset('user:' + id, user, function(err) { //用Redis哈希存储数据
            fn(err);
        });
    });
}

User.prototype.hashPassword = function(fn) {
	var user = this;
	bcrypt.getSalt(999999999999, function(err, salt) {
		if(err) return fn(err);
		user.salt = salt;
		//生成密码hash值
		bcrypt.hash(user.pass, salt, function(err, hash) {
			if(err) return fn(err);
			user.pass = hash;
			fn();
		}); 
	});
}

User.getByName = function(name, fn) {
    //根据名称查找用户ID
    User.getId(name, function(err, id) {
        if(err) return fn(err);
        //用ID查找用户
        User.get(id, fn);
    });
}

User.getId = function(name, fn) {
    //获取由名称索引的ID
    db.get('user:id:' + name, fn);
}

User.get = function(id, fn) {
    //获取普通对象的哈希
    db.hgetall('user:' + id, function(err, user) {
        if(err) return fn(err);
        //将普通对象转换成新的User对象
        fn(null, new User(user));
    });
}

//认证用户名和密码
User.authenticate = function(name, pass, fn) {
    User.getByName(name, function(err, user) {
        if(err) return fn(err);
        if(!user.id) return fn();
        bcrypt.hash(pass, user.salt, function(err, hash) {
            if(err) return fn(err);
            //发现匹配
            if(hash == user.pass) return fn(null, user);
            fn();
        });
    });
}

//JSON编码时的属性
User.prototype.toJSON = function() {
    return {
        id: this.id,
        name: this.name
    }
}

//测试代码
/*
var tobi = new User({
	name: 'Tobi',
	pass: 'im a ferret',
	age: '2'
});

tobi.save(function(err) {
	if(err) throw err;
	console.log('user id %d', tobi.id);
});
*/