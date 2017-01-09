var http = require('http')
var cheerio = require('cheerio')
var url = 'http://127.0.0.1/'

function filterChapters(html) {
	var $ = cheerio.log(html)
	var chapters = $('.lernchapter')

	var courseData = []

	chapters.each(function(item) {
		var chapter = $(this)
		var chapterTitle = chapter.find('strong').text()
	})
}

http.get(url, function(res) {
	var html = ''
	
	//数据载入事件
	res.on('data', function(data) {
		html += data
	})

	//数据载入结束
	res.on('end', function() {
		filterChapters(html)
	})
}).on('error', function() {
	console.log('获取课程数据出错!')
})