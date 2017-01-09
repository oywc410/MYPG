var redis = require('redis');

//连接服务器 redis
var client = redis.createClient(6379, '127.0.0.1');

client.on('error', function(err) {
    console.log('Error ' + err);
});

//数据操作
client.set('color', 'red', redis.print);
client.get('color', function(err, value) {
    if (err) throw err;
    console.log('Got: ' + value);
});

//用哈希表储存和获取数据
client.hmset('camping', {
    'shelter': '2-person tent',
    'cooking': 'campstove',
}, redis.print);

//获取元素"cooking"的值
client.hget('camping', 'cooking', function(err, value) {
    if (err) throw err;
    console.log('Will be connking with: ' + value);
});

//获取哈希表的键
client.hkeys('camping', function(err, keys) {
    if (err) throw err;
    keys.forEach(function(key, i) {
        console.log(' ' + key);
    });
});

//用链表储存和获取数据
client.lpush('tasks', 'Paint the bikeshed red.', redis.print);
client.lpush('tasks', 'Paint the bikeshed green.', redis.print);
client.lrange('tasks', 0, -1, function(err, items) {
    if (err) throw err;
    items.forEach(function(item, i) {
        console.log(' ' + item);
    });
});

//用集合储存和获取数据
client.sadd('ip_addresses', '204.10.37.96', redis.print);
client.sadd('ip_addresses', '204.10.37.96', redis.print);
client.sadd('ip_addresses', '72.32.231.8', redis.print);
client.smembers('ip_addresses', function(err, members) {
    if (err) throw err;
    console.log(members);
});
