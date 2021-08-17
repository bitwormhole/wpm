// vue.config.js
module.exports = {
    configureWebpack: {
        plugins: [
            //     new MyAwesomeWebpackPlugin()
        ],

    },
    chainWebpack: config => {
        config
            .plugin('html')
            .tap(args => {
                args[0].title = 'WPM Station'
                return args
            })
    },

    devServer: {
        open: false,
        port: 18080,
        proxy: {
            '/api/': {
                target: 'http://localhost:18081/',
                changeOrigin: true
            },
        }
    }

}