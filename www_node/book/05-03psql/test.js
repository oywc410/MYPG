var pg = require('pg')
var conString = "tcp://postgres:root@127.0.0.1/node_psql_test";

var client = new pg.Client(conString);
client.connect();

//插入数据
client.query(
	'INSERT INTO user ' +
	"(name) VALUES ('Mike')"
);

//占位符形式插入数据
client.query(
	"INSERT INTO users " +
	"(name, age) VALUES ($1, $2)",
	['Mike', 39]
);

//要在插入一条记录后得到它的主键值，可以用RETURNING从句加上列名指定想要返回哪一列的值
client.query(
	"INSERT INTO users " +
	"(name, age) VALUES($1, $2)" + 
	"RETURNING id",
	['Mike', 39],
	function(err, result) {
		if(err) throw err;
		console.log('Insert ID is ' + result.rows[0].id);
	}
);

//查询
var query = client.query(
	"SELECT * FROM users WHERE age > $1",
	[40]
);

//处理查询返回的记录
query.on('row', function(row) {
	console.log(row.name);
});

//查询完成后的处理
query.on('end', function() {
	client.end();
});