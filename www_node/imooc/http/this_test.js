var pet = {
	words: '...',
	speak: function() {
		//this指向pet
		console.log(this.words)
		console.log(this === pet)
	}
}

pet.speak()

function pet2(words) {
	this.words = words
	//this指向global 顶端
	console.log(this.words)
	console.log(this === global)
}

pet2('aaa');

function Pet3(words) {
	this.words = words
	this.speak = function() {
		//指向构造体 class
		console.log(this.words)
		console.log(this)
	}
}

var pet3 = new Pet3('...');
pet3.speak();