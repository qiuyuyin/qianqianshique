const { mergeSassVariables } = require('@vuetify/cli-plugin-utils')

module.exports = {
  publicPath: process.env.NODE_ENV === 'production'
    ? '/'
    : '/',
  devServer: {
    open: false, // 是否打开浏览器
    host: 'localhost',
    port: '8080', // 代理断就
    https: false,

    // hotOnly: false, // 热更新
    proxy: {
      '/token/': {
        target:
          'https://openai.wps.cn', // 目标代理接口地址
        secure: false,
        changeOrigin: true, // 开启代理，在本地创建一个虚拟服务端
        ws: true, // 是否启用websockets
        pathRewrite: {
          '^/token': '/',
        },
      },
      '/server/': {
        target:
          'http://47.110.147.51:8888', // 目标代理接口地址
        secure: false,
        changeOrigin: true, // 开启代理，在本地创建一个虚拟服务端
        ws: true, // 是否启用websockets
        pathRewrite: {
          '^/server': '/',
        },
      },
    },
    headers: {
      'Access-Control-Allow-Origin': '*',
    },
  },
  transpileDependencies: ['vuetify'],
  chainWebpack: config => {
    const modules = ['vue-modules', 'vue', 'normal-modules', 'normal']
    modules.forEach(match => {
      config.module
        .rule('sass')
        .oneOf(match)
        .use('sass-loader')
        .tap(opt => mergeSassVariables(opt, "'@/styles/variables.scss'"))
      config.module
        .rule('scss')
        .oneOf(match)
        .use('sass-loader')
        .tap(opt => mergeSassVariables(opt, "'@/styles/variables.scss';"))
    })
  },
}
