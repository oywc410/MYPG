var EventEmitter = require('events').EventEmitter

var life = new EventEmitter()

//设置最大值
life.setMaxListeners(11)

//addEventListener
//注册事件(设置监听器,最好不要超过10个)
life.on('事件名', function(who) {
	console.log('...' + who);
})

life.on('事件名', function(who) {
	console.log('...' + who);
})

life.on('事件名', function(who) {
	console.log('...' + who);
})

life.on('事件名', function(who) {
	console.log('...' + who);
})

function water() {
	console.log('监听移除')
}

life.on('事件名', water)

life.once('事件名', water)//响应只应该发生一次的事件

//移除事件监听 秘名函数无效
life.removeListener('事件名', water)
//移除全部监听
life.removeAllListeners()
life.removeAllListeners('事件名')

//获取事件监听数量
console.log(life.listeners('事件名').length)
console.log(EventEmitter.listenerCount(life, '事件名'))

//触发事件
var hasConfortListener = life.emit('事件名', 'a')
//hasConfortListener 判断事件是否触发