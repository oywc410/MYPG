var fs = require('fs');
var path = require('path');
var args = process.argv.splice(2);//获取输入参数
var command = args.shift();

var taskDescription = args.join(' ');//合并剩余的参数

//process.cmd() 当前执行命令的目录
var file = path.join(process.cwd(), '/.tasks')//根据当前的工作目录解析数据库的相对路径

//判断命令
switch(command) {
	case 'list':
		listTasks(file);
		break;
	case 'add':
		addTask(file, taskDescription);
		break;
	default:
		console.log('Usage: ' + process.argv[0] + ' list|add [taskDescription]');
}

//从文本文件中加载JSON编码数据
function loadOrinitializeTaskArray(file, cb) {
	fs.exists(file, function(exists) {
		var tasks = [];
		if(exists) {
			fs.readFile(file, 'utf8', function(err, data) {
				if(err) throw err;
				var data = data.toString();
				var tasks = JSON.parse(data || '[]');
				cb(tasks);
			});
		} else {
			cb([]);
		}
	});
}

//列出任务的函数
function listTasks(file) {
	loadOrinitializeTaskArray(file, function(tasks) {
		for(var i in tasks) {
			console.log(tasks[i]);
		}
	});
}

//把任务保持到磁盘中
function storeTasks(file, tasks) {
	fs.writeFile(file, JSON.stringify(tasks), 'utf8', function(err) {
		if(err) throw err;
		console.log('Saved.');
	});
}

//添加一个任务
function addTask(file, taskDescription) {
	loadOrinitializeTaskArray(file, function(tasks) {
		tasks.push(taskDescription);
		storeTasks(file, tasks);
	});
}


/*
	cmd:
		node cli_task.js add Floss the cat.
		node cli_task.js list
		node cli_task.js add Buy some hats.
*/