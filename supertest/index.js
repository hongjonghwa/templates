const request = require('supertest')
const assert = require('assert')

request('https://m.naver.com/')
  .get('/')
  //.expect('Content-Type', /json/)
  //.expect('Content-Length', '15')
  .expect(200)
  .timeout(1000)
  .end(function (err, res) {
    if (err) throw err
    assert.ok(res)
  })
