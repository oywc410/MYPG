var http = require('http')
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