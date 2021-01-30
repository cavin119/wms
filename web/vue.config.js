const path = require('path')

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
    //end configureWebpack
    chainWebpack(config) {
        config.plugin('html').tap(args => {
            args[0].title = "仓储管理系统-WMS"
            return args
        })
    }
    //end chainWebpack
}