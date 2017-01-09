var http = require('http');

var server = http.createServer(function(req, res) {

    req.setEncoding('utf8'); //指定编码后 chunk 就为字符串数据
    req.on('data', function(chunk) { //只要读入了新的数据块，就触发data事件
        //默认 chunk 为buffer数据
        console.log('parsed', chunk);
    });

    req.on('end', function() { //数据全部读完之后触发
        console.log('done parsing');

        var body = 'test';
        //sres.setHeader('Location', 'http://127.0.0.1');
        res.setHeader('Content-length', body.length);
        res.setHeader('Content-length', 'text/html');
        res.statusCode = 200;
        res.end(body);
    });
});

server.listen(3000, "127.0.0.1");;
