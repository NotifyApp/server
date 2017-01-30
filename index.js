const server = require('http').createServer()
const url = require('url')
const WebSocketServer = require('ws').Server
const wss = new WebSocketServer({ server: server })
const express = require('express')
const app = express()
const bodyParser = require('body-parser')
const port = 5000

app.use(bodyParser.json())

wss.on('connection', (ws) => {
  ws.send(JSON.stringify({
    message: 'Hello'
  }))
})

app.get('/', (req, res) => {
  res.json({
    success: true,
    message: 'Notify API',
    version: "1.0.0"
  })
})

app.post('/notify', (req, res) => {
  if (!req.body.title || req.body.title === '' || !req.body.text || req.body.text === '') {
    return res.json({
      success: false,
      message: 'Fields empty'
    })
  }
  wss.clients.forEach((ws) => {
    ws.send(JSON.stringify({
      title: req.body.title,
      text: req.body.text
    }))
  })
  res.json({
    success: true,
    message: 'Sent!'
  })
})

server.on('request', app);
server.listen(port, () => {
  console.log('Listening on ' + port)
})