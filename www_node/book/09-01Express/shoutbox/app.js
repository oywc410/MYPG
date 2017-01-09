var express = require('express');
var path = require('path');
var favicon = require('serve-favicon');
var logger = require('morgan');
var cookieParser = require('cookie-parser');
// 首先引入 express-session 这个模块
var session = require('express-session');
var bodyParser = require('body-parser');

var routes = require('./routes/index');
var users = require('./routes/users');

//引入中间件
var user = require('./lib/middleware/user');
var validate = require('./lib/middleware/validate');
var page = require('./lib/middleware/page');

var Entry = require('./lib/entry');

//引入路由逻辑
var register = require('./routes/register');
var login = require('./routes/login');
var entries = require('./routes/entries');
var api = require('./routes/api');

//信息提示处理
var messages = require('./lib/messages');

var app = express();

// view engine setup
app.set('views', path.join(__dirname, 'views'));
app.set('view engine', 'ejs');

// uncomment after placing your favicon in /public
//app.use(favicon(path.join(__dirname, 'public', 'favicon.ico')));
app.use(logger('dev'));
app.use(bodyParser.json({
    limit: '1mb'
}));
app.use(bodyParser.urlencoded({
    extended: true
}));
app.use(cookieParser());
// 按照上面的解释，设置 session 的可选参数
app.use(session({
    secret: 'keyboard cat',
    cookie: {
        maxAge: 60000
    }
}));
app.use(messages);
app.use(user);
app.use(express.static(path.join(__dirname, 'public')));

//app.use('/', routes);
//app.use('/users', users);

//添加路由
app.get('/register', register.form);
app.post('/register', register.submit);
app.get('/login', login.form);
app.post('/login', login.submit);
app.get('/logout', login.logout);
//app.get('/', entries.list);
app.get('/:page?', page(Entry.count, 5), entries.list);
app.get('/post', entries.form);
app.post('/post', 
          validate.required('entry[title]'),
          validate.lengthAbove('entry[title]', 4),
          entries.submit);

//app.use('/api', api.auth);
app.get('/api/user/:id', api.user);
app.get('/api/entries/:page?', page(Entry.count),api.entries);
app.post('/api/entry', entries.submit);

//添加自定义404页面
app.use(routes.notfound);

// catch 404 and forward to error handler
app.use(function(req, res, next) {
    var err = new Error('Not Found');
    err.status = 404;
    next(err);
});

// error handlers

// development error handler
// will print stacktrace
if (app.get('env') === 'development') {
    app.use(function(err, req, res, next) {
        res.status(err.status || 500);
        res.render('error', {
            message: err.message,
            error: err
        });
    });
}

// production error handler
// no stacktraces leaked to user
app.use(function(err, req, res, next) {
    res.status(err.status || 500);
    res.render('error', {
        message: err.message,
        error: {}
    });
});


module.exports = app;
