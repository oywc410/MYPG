var flow = require('nimble');
var exec = require('child_process').exec;

function downloadNodeVersion(version, destination, callback) {
    var url = 'http://nodejs.org/dist/node-v' + version + '.tar.gz';
    var filepath = destination + '/' + version + '.tgz';
    exec('curl ' + url + ' >' + filepath, callback);
}

flow.series([ //按顺序执行串行任务
    function(callback) {
        flow.parallel([ //并行下载
            function(callback) {
                console.log('Downloading Node v0.4.6...');
                downloadNodeVersion('0.4.6', './tmp', callback);
            },
            function(callback) {
                console.log('Downloading Node v0.4.7...');
                downloadNodeVersion('0.4.7', './tmp', callback);
            },
            function(callback) {
            	flow.series([//并行中执行串行下载
            		function(callback) {
            			console.log('Downloading Node v0.4.8...');
            			downloadNodeVersion('0.4.8', './tmp', callback);
            		},
            		function(callback) {
            			console.log('Downloading Node v0.4.9...');
            			downloadNodeVersion('0.4.9', './tmp', callback);
            		}
            	], callback);
            },
            function(callback) {
                console.log('Downloading Node v0.5.0...');
                downloadNodeVersion('0.5.0', './tmp', callback);
            }
        ], callback)
    },
    function(callback) {
        console.log('Creating archive of dowmloaded file...');
        exec(
            'tar cvf node_distros.tar ./tmp/0.4.6.tgz ./tmp/0.4.7.tgz',
            function(error, stdout, stderr) {
                console.log('All done!');
                callback();
            }
        );
    }
    ]
);
