var connect = require('connect')
var http = require('http')
var vhost = require('vhost')
 
// create main app 
var app = connect()
 
app.use(vhost('mail.example.com', function (req, res) {
  // handle req + res belonging to mail.example.com 
  res.setHeader('Content-Type', 'text/plain')
  res.end('hello from mail!')
}))
 
// an external api server in any framework 
var httpServer = http.createServer(function (req, res) {
  res.setHeader('Content-Type', 'text/plain')
  res.end('hello from the api!')
})
 
app.use(vhost('api.example.com', function (req, res) {
  // handle req + res belonging to api.example.com 
  // pass the request to a standard Node.js HTTP server 
  httpServer.emit('request', req, res)
}))
 
app.listen(3000)