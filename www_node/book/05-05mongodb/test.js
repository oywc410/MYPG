var mongodb = require('mongodb');
var server = new mongodb.Server('127.0.0.1', 27017, {});

var client = new mongodb.Db('mydatabase', server, {w: 1});

client.open(function(err) {
	if(err) throw err;
	client.collection('test_insert', function(err, collection) {
		if(err) throw err;
		console.log('把mongoDB查询代码放在这里');

		//追加文档
		collection.insert(
			{
				"title": "I like cake",
				"body": "It is quite good."
			},
			{safe: true},//在查询语句中声明{safe: true}表明你想让数据库操作在执行回调之前完成。
			function(err, documents) {
				if(err) throw err;
				console.log('Document ID is: ' + documents[0].id);
			}
		);

		//更新文档
		var _id = new client.bson_serializer.ObjectID('xxxx');//documents[0].id

		collection.update(
			{_id: _id},
			{$set: {"title": "I ate too much cake"}},
			{safe: true},
			function(err) {
				if(err) throw err;
			}
		);

		//搜索文档
		collection.find({"title": "I like cake"}).toArray(
			function(err, results) {
				if(err) throw err;
				console.log(results);
			}
		);

		//删除文档
		var _id = new client.bson_serializer.ObjectID('xxxx');//documents[0].id
		collection.remove({_id: _id}, {safe: true}, function(err) {
			if(err) throw err;
		});


		client.close();
	});
});

