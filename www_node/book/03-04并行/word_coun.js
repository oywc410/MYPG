//并行化流程控制 检索文本关键字次数
var fs = require('fs');
var completedTaaks = 0;
var tasks = [];
var wordCounts = {};
var filesDir = './text';

//当所有任务全部完成后,列出文件中用的单词以及个数
function checkIfComplete() {
	completedTaaks++;
	if(completedTaaks == tasks.length) {
		for(var index in wordCounts) {
			console.log(index + ': ' + wordCounts[index]);
		}
	}
}

//统计文本中出现的单词数
function countWordsInText(text) {
	var words = text.toString().toLowerCase().split(/\W+/).sort();
	for(var index in words) {
		var word = words[index];
		if(word) {
			wordCounts[word] = (wordCounts[word]) ? wordCounts[word] + 1 : 1;
		}
	}
}

fs.readdir(filesDir, function(err, files) {
	if(err) throw err;
	for(var index in files) {
		var task = (function(file) {
			return function() {
				fs.readFile(file, function(err, text) {
					if(err) throw err;
					countWordsInText(text);
					checkIfComplete();
				});
			}
		})(filesDir+ '/' + files[index]);
		tasks.push(task);
	}

	for(var task in tasks) {
		tasks[task]();
	}
});