var express = require('express');

module.exports = function(req, res, next) {
	//locals 可以在视图中任意使用messages,removeMessages
	res.locals.messages = req.session.messages || [];
	if(!req.session.t) {
		req.session.t = 1;
	} else {
		req.session.t++;
	}
	
	res.locals.removeMessages = function() {
		//req.session.messages = [];
	};

	res.message = function(msg, type) {
		type = type || 'info';
		var sess = req.session;
		sess.messages = sess.messages || [];
		sess.messages.push({type: type, string: msg});
	}

	res.error = function(msg) {
		return res.message(msg, 'error');
	}

	next();
}