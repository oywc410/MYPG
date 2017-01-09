var http = require('http');
var parse = require('url').parse;
var join = require('path').join;
var fs = require('fs');

var root = __dirname; //当前文件目录

var server = http.createServer(function(req, res) {
    var url = parse(req.url);
    var path = join(root, url.pathname); //构建绝对路径
    console.log(path);

    fs.stat(path, function(err, stat) { //文件预加载系统
        if (err) {
            if ('ENOENT' == err.code) {
                res.statusCode = 404;
                res.end('Not Found');
            } else {
                res.statusCode = 500;
                res.end('Internal Server Error');
            }
        } else {
            res.setHeader('Content-Length', stat.size);
            var stream = fs.createReadStream(path);

            stream.pipe(res); //res.end(); 将在pipe中被调用

            stream.on('error', function(err) {
                console.log(err);
                res.statusCode = 500;
                res.end('Internal Server Error');
            });
            /*
		    stream.on('data', function(chunk) {
		        res.write(chunk);
		    });

		    stream.on('end', function() {
		        res.end();
		    });
			*/
        }
    });


});

server.listen(3000, "127.0.0.1");


// 客户端访问 curl http://127.0.0.1:3000/test.js