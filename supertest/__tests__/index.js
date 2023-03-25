const request = require('supertest')
const assert = require('assert')

// USAGE: https://www.npmjs.com/package/supertest

describe('Hi Naver', () => {
  it('hahaha', (done) => {
    expect.assertions(1)
    request('https://m.naver.com/')
      .get('/')
      //.expect('Content-Type', /json/)
      //.expect('Content-Length', '15')
      .expect(200)
      .timeout(10)
      .end(function (err, res) {
        if (err) done(err)
        expect(res).toBeDefined() //expect.assertions 1
        done()
      })
  })

  it('hello', (done) => {
    request('https://m.naver.com/')
      .get('/')
      //.expect('Content-Type', /json/)
      //.expect('Content-Length', '15')
      .timeout(1000)
      .expect(200, done)
  })
})
