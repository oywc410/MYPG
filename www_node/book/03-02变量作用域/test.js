function asyncFunction(callback) {
	setTimeout(callback, 200)
}

var color = 'blue';

asyncFunction(function() {
	//输出the color is qreen
	console.log('The color is ' + color);
})

var echolog = function(color) {
	asyncFunction(function() {
		//输出the color is blue (变量的作用域不同)
		console.log('The color is ' + color);
	})
};

echolog(color);

color = 'qreen';

