'use strict'
const merge = require('webpack-merge')
const prodEnv = require('./prod.env')

module.exports = merge(prodEnv, {
  NODE_ENV: '"development"',
  // BASE_API: '"http://120.27.63.9:8080"'
  // BASE_API:'"http://120.24.221.188:40001/mock/11/"'
  // BASE_API : '"http://192.168.0.103:8000/"'
  //API网关
  BASE_API : '"http://172.29.133.152:8000/"'
})
