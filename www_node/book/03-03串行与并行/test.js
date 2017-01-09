//串行化运行控制
var flow = require('nimble');

//串行化运行
flow.series([
	function (callback) {
		setTimeout(function() {
			console.log('I execut first.');
			callback();
		}, 1000);
	},
	function (callback) {
		setTimeout(function() {
			console.log('I execut next.');
			callback();
		}, 500);
	},
	function (callback) {
		setTimeout(function() {
			console.log('I execut last.');
			callback();
		}, 100)
	}
]);