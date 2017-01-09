//网站小图标
//日志
var connect = require('connect');
var favicons = require('connect-favicons');

var app = connect()
			.use(favicons(__dirname + '/public/icos'))
			/**
			favicon.ico
			apple-touch-icon.png
			apple-touch-icon-precomposed.png
			apple-touch-icon-57x57.png
			apple-touch-icon-57x57-precomposed.png
			apple-touch-icon-72x72.png
			apple-touch-icon-72x72-precomposed.png
			apple-touch-icon-114x114.png
			apple-touch-icon-114x114-precomposed.png
			apple-touch-icon-144x144.png
			apple-touch-icon-144x144-precomposed.png
			*/
			.use(function(req, res, next) {
				res.end('logs');
			}).listen(3000);