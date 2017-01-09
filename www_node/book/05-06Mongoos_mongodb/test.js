var mongoose = require('mongoose');
var db = mongoose.connect('mongodb://localhost/tasks');

//注册Schema
var Schema = mongoose.Scheme;
var Tasks = new Schema({
	project: String,
	description: String
});
mongoos.model('Task', Tasks);

//添加任务
var Task = mongoose.model('Task');
var task = new Task();
task.project = 'Bikeshed';
task.description = 'Paint the bikeshed red.';
task.save(function(err) {
	if(err) throw err;
	console.log('Task saved.');
});

//搜索文档
var Task = mongoose.model('Task');
Task.find({'project': 'Bikeshed'}, function(err, tasks) {
	for(var i = 0; i < tasks.length; i++) {
		console.log('ID: ' + tasks[i]._id);
		console.log(tasks[i].description);
	}
});

//更新文档
var Task = mongoose.model('Task');
Task.update(
	{_id: 'xxxxx'},
	{description: 'aaaa'},
	{multi: false}, //只更新一个文档
	function(err, rows_updated) {
		if(err) throw err;
		console.log('Updated.');
	}
);

//删除文档
var Task = mongoose.model('Task');
Task.findById('xxxx', function(err, task) {
	task.remove();
});

mongoose.disconnect();终止连接