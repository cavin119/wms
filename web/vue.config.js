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
            // 把key的路径代理到target位置
            // detail: https://cli.vuejs.org/config/#devserver-proxy
            // [process.env.VUE_APP_BASE_API]: { //需要代理的路径   例如 '/api'
            //     target: `http://127.0.0.1:8888/`, //代理到 目标路径
            //     changeOrigin: true,
            //     pathRewrite: { // 修改路径数据
            //         ['^' + process.env.VUE_APP_BASE_API]: '' // 举例 '^/api:""' 把路径中的/api字符串删除
            //     }
            // }
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