var mongoose = require('mongoose');
mongoose.connect('mongodb://127.0.0.1/photo_app');

module.exports = mongoose.model('Photo', {
    name: String,
    path: String
});
