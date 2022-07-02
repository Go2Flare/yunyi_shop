var express = require("express");
var webpackDevMiddleware = require("webpack-dev-middleware");
var webpackHotMiddleware = require("webpack-hot-middleware");
var webpack = require("webpack");
var webpackConfig = require("./webpack.config");

var app = express();
var compiler = webpack(webpackConfig);

app.use(webpackDevMiddleware(compiler, {
	publicPath: 'http://120.24.221.188:3000/',
	index: 'index.html'
}))

app.use(webpackHotMiddleware(compiler, {
    log: console.log
}))

app.listen(3000, function() {
	console.log("listening on port 3000");
})
