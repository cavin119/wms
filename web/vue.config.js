const path = require('path')

function resolve(dir) {
    return path.join(__dirname, dir)
}

module.exports = {
    devServer: {
        port: 8080,
        open: true,
        overlay: {
            warnings: false,
            errors: true
        },
        proxy: {
            '^/api': { //需要代理的路径   例如 '/api'
                target: `http://127.0.0.1:8888/`, //代理到 目标路径
                changeOrigin: true
            }
        },
    },
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