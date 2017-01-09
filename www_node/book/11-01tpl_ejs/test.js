var ejs = require('ejs');
var template = '<%= message %>\n'; //=转义输出
template += '<%- message %>'; //-非转义输出
var context = {
    message: 'Hello template!<script>alert("alert!");</script>'
};

//自定义标签
ejs.open = "<%";
ejs.close = "%>";

console.log(ejs.render(template, context));
