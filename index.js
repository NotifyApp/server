const server = require('http').createServer()
const url = require('url')
const WebSocketServer = require('ws').Server
const wss = new WebSocketServer({ server: server })
const express = require('express')
const app = express()
const port = 5000
 
app.use((req, res) => {
  res.send({ msg: "hello" })
})
 
wss.on('connection', (ws) => {
  ws.on('message', (message) => {
    console.log('received: %s', message)
  })

  ws.send('something')
})
 
server.on('request', app);
server.listen(port, () => {
  console.log('Listening on ' + port)
})