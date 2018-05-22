# go_hprose
go_hprose 通信demo

nodejs端
```js
const hprose = require('hprose')
const client = hprose.Client.create('tcp4://127.0.0.1:8888/')
client.subscribe('push', 'id', function (date) {
  console.log(date)
})
```

[hprose](https://github.com/hprose/hprose-nodejs/wiki/%E6%8E%A8%E9%80%81%E6%9C%8D%E5%8A%A1)