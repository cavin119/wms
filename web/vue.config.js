'use strict'

const path = require('path')
const packageConf = require('./package.json')

function resolve(dir) {
    return path.join(__dirname, dir)
}

module.exports = {
    configureWebpack: {
        //@路径走src文件夹
        resolve: {
            alias: {
                '@': resolve('src')
            }
        }
    },
    chainWebpack(config) {
        config.plugin('html').tap(args => {
            args[0].title = "仓储管理系统-WMS"
            return args
        })
    }
}