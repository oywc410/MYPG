var http = require('http')
var querystring = require('querystring')

var postData = querystring.stringify({
	'content': 'aaa',
	'cid' : '123'	
})

var options = {
	hostname : '127.0.0.1',
	port: 80,
	path: '/course/docomment',
	method: 'POST',
	headers: {//送信头
		'Accept' : 'aaa',
		'Cookie' : 'bbbb',
		'Content-Length' : postData.length
	}
}

var req = http.request(options, function(res) {
	console.log('Status:' + res.statusCode)
	console.log('headers:' + JSON.stringify(res.headers))

	//注册事件
	res.on('data', function (chunk) {
		console.log(Buffer.isBuffer(chunk))
		console.log(typeof chunk)
	})

	res.on('end', function() {
		console.log('评论完毕')
	})

	res.on('error', function(e) {
		console.log('Error' + e.message)
	})
})

//写入发送信息内容
req.write(postData)
req.end()



var url = 'http://127.0.0.1/'

http.get(url, function(res) {
	var html = ''
	
	//数据载入事件
	res.on('data', function(data) {
		html += data
	})

	//数据载入结束
	res.on('end', function() {
		console.log(html)
	})
}).on('error', function() {
	console.log('获取课程数据出错!')
})