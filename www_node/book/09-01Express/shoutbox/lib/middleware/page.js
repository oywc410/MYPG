module.exports = function(fn, perpage) {
    //每页记录条数的默认值
    perpage = perpage || 10;
    //返回中间件函数
    return function(req, res, next) {
        var page = Math.max(
            parseInt(req.param('page') || '1', 10),
            1
        ) - 1;
        //fn 获取条数的回调函数
        fn(function(err, total) {
            if (err) return next(err);

            //保存page属性
            req.page = res.locals.page = {
                number: page,
                perpage: perpage,
                from: page * perpage,
                to: page * perpage + perpage - 1,
                total: total,
                count: Math.ceil(total / perpage)
            };

            next();
        });
    }
}
