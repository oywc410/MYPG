var mongoose = require('mongoose');
mongoose.connect('mongodb://localhost/tasks');

var Cat = mongoose.model('Cat', {
    name: String
});

var kitty = new Cat({
    name: 'zzz'
});
kitty.save(function(err, aaa) {
    console.log(err);
    console.log(aaa);
});


var Task = mongoose.model('Cat');
Task.find({
    name: 'zzz'
}, function(err, tasks) {
    console.log(err);
    console.log(tasks);

});
