//RSS读取器

var fs = require('fs');
//http信息
var request = require('request');
//html解析为js对象
var htmlparser = require('htmlparser');
var configFilename = './rss_feeds.txt';

//利用串行话工具获取网站的RSS

function checkForRSSFile() {
    fs.exists(configFilename, function(exists) {
        if (!exists) {
            return next(new Error('Missing RSS file: ' + configFilename));
        }
        next(null, configFilename);
    });
}

function readRSSFile(configFilename) {
    //读取并解析包含预订源URL的文件
    fs.readFile(configFilename, function(err, feedList) {
        if (err) return next(err);

        feedList = feedList.toString().replace(/^\s+|\s+$/g, '').split("\n");
        //从预订源URL数组中随机选择一个预订源URL
        var random = Math.floor(Math.random() * feedList.length);
        next(null, feedList[random]);
    });
}

function downloadRSSFeed(feedUrl) {
    //发送http请求
    request({
        uri: feedUrl
    }, function(err, res, body) {
        if (err) return next(err);
        if (res.statusCode != 200) {
            return next(new Error('Abnormal response status code'));
        }
        next(null, body);
    });
}

function parseRSSFeed(rss) {
    var handler = new htmlparser.RssHandler();
    var parser = new htmlparser.Parser(handler);
    parser.parseComplete(rss);
    if (!handler.dom.items.langth) {
        return next(new Error('NO RSS items found'));
    }

    var item = handler.dom.items.shift();
    console.log(item.title);
    console.log(item.link);
}

//把要做的任务按执行顺序到一个数组中
var tasks = [checkForRSSFile, readRSSFile, downloadRSSFeed, parseRSSFeed];

//负责执行任务
function next(err, result) {
    if (err) throw err; //如果任务出错,则抛出异常
    //从任务数字中取出下个任务
    var currenTask = tasks.shift();

    if (currenTask) {
        currenTask(result); //执行当前任务
    }
}

next(); //开始任务的串行化执行
