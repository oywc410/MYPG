var http = require('http'); //HTTP服务器和客户端功能
var fs = require('fs'); //文件系统路径相关功能
var path = require('path'); //HTTP服务器和客户端功能
var mime = require('mime'); //文件名扩展得出MIME类型的能力

var chatServer = require('./lib/chat_server');

var cache = {} //cache使用来缓存文件内容的对象

//发送404错误
function send404(response) {
    response.writeHead(404, {
        'Content-Type': 'text/plain'
    });
    response.write('Error 404: resource not found.');
    response.end();
}

//发送文件内容
function sendFile(response, filePath, fileContents) {
    response.writeHead(200, {
        "Content-Type": mime.lookup(path.basename(filePath))
    })
    response.end(fileContents);
}

//提供静态文件服务
function serveStatic(response, cache, absPath) {
    if (cache[absPath]) {
        sendFile(response, absPath, cache[absPath]); //从内存中返回文件
    } else {
        fs.exists(absPath, function(exists) { //检查文件是否存在
            if (exists) {
                //从硬盘中读取文件
                fs.readFile(absPath, function(err, data) {
                    if (err) {
                        send404(response);
                    } else {
                        cache[absPath] = data;
                        //从硬盘中读取文件内容并返回
                        sendFile(response, absPath, data);
                    }
                })
            } else {
                send404(response);
            }
        })
    }
}

var server = http.createServer(function(request, response) {
    var filePath = false;

    if (request.url == '/') {
        filePath = 'public/index.html';
    } else {
        filePath = 'public' + request.url;
    }

    var absPath = './' + filePath;
    serveStatic(response, cache, absPath);
})

chatServer.listen(server);

server.listen(3000, function() {
    console.log("Server lostening on port 3000.");
});
