var redis = require('redis');
var db = redis.createClient();

//向模块中输出Entry函数
module.exports = Entry;

function Entry(obj) {
    for (var key in obj) {
        this[key] = obj[key];
    }
}

//new 时可使用save方法    var entry = new Entry; entry.save;
Entry.prototype.save = function(fn) {
    var entryJSON = JSON.stringify(this);

    //将JSON字符串保持到Redis列表中
    db.lpush(
        'entries',
        entryJSON,
        function(err) {
            if (err) return fn(err);
            fn();
        }
    );
}

Entry.getRange = function(from, to, fn) {
    //获取entries消息记录
    db.lrange('entries', from, to, function(err, items) {
        if (err) return fn(err);
        var entries = [];

        //解码之前保持的JSON消息记录
        items.forEach(function(item) {
            entries.push(JSON.parse(item));
        });

        fn(null, entries);
    });
}

Entry.count = function(fn) {
    db.llen('entries', fn);
}
