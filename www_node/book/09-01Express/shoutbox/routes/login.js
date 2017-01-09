var User = require('../lib/user');

exports.form = function(req, res) {
    console.log('aaaa');
    res.render('login', {
        title: 'Login'
    });
}

exports.submit = function(req, res, next) {
    var data = req.body.user;
    User.authenticate(data.name, data.pass, function(err, user) {
        if (err) return next(err);
        if (user) {
            req.session.uid = user.id;
            res.redirect('/');
        } else {
            //输入错误信息
            res.error("Sorry! invalid credentials.");
            res.redirect('back');
        }
    });
}

exports.logout = function(req, res) {
    req.session.destroy(function(err) {
        if (err) throw err;
        res.redirect('/');
    });
}
