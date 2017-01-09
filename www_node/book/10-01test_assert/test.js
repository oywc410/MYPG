var assert = require('assert');
var Todo = require('./todo');

var todo = new Todo();
var testsCompleted = 0;

function deleteTest() {
	todo.add('Dlelte Me');
	//断言 equal =
	assert.equal(todo.getCount(), 1, '1 item should exist');
	todo.deleteAll();
	//断言
	assert.equal(todo.getCount(), 0, 'No items shold exist');
	//记录测试完成
	testsCompleted++;
}

function addTest() {
	todo.deleteAll();
	todo.add('Added');
	//断言
	assert.notEqual(todo.getCount(), 0, '1 item should exits');
	testsCompleted++;
}

//equal notEqual strictEqual notStrictEqual  ==  !=  === !== 对比值
//deepEqual notDeepEqual 对比对象

//检查回调函数
function doAsyncTest(cb) {
	todo.doAsync(function(value) {
		//断言值为true
		assert.ok(value, 'Callback should be passed true');
		testsCompleted++;
		cb();
	});
}

//检查是否抛出异常
function throwsTest(cb) {
	//不带参数调用todo.add 第二个参数为查找是否存在require
	assert.throws(todo.add, /require/);
	testsCompleted++;
}


deleteTest();
addTest();
throwsTest();
doAsyncTest(function() {
	//表面结束测试
	console.log('Completed ' + testsCompleted + ' tests');
});