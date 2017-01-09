var Photo = require('../models/Photo');
var path = require('path');
var fs = require('fs');
var join = path.join;

//路由列表
exports.list = function(req, res) {

	Photo.find({}, function(err, photos) {
		res.render('photos', {
			title: 'Photos',
			photos: photos
		});
	});
}

//上传表单
exports.form = function(req, res) {
    res.render('photos/upload', {
        title: 'Photo upload'
    });
}

//表单提交处理
exports.submit = function(dir) {
    return function(req, res, next) {
        console.log(req.files);
        var img = req.files.photo.image;
        //默认为原来的文件名
        var name = req.body.photo.name || img.name;
        var path = join(dir, img.name);
        console.log(path);

        //重命名文件
        fs.rename(img.path, path, function(err) {
            if (err) return next(err);

            var data = new Photo({
                name: name,
                path: img.name
            });
            //将数据保存到数据库
            data.save(function(err, aaa) {
                //委派错误,重定向到首页
                if (err) return next(err);
                res.redirect('/');
            });
        });
    }
}

//文件下载
exports.download = function(dir) {
	return function(req, res, next) {
		var id = req.params.id;
		Photo.findById(id, function(err, photo) {
			if(err) {
				return next(err);
			}
			var path = join(dir, photo.path);//构建文件的绝对路径
			res.sendfile(path);//文件传输
			//res.download(path, photo.name + '.jpeg');//下载文件
		});
	}
}
