const { DefinePlugin } = require('webpack')

module.exports = {
  configureWebpack: {
    plugins: process.env.NODE_ENV === 'production'
      ? [
        new DefinePlugin({
          __VERSION__: JSON.stringify(require('./package.json').version)
        })
      ]
      : [
        new DefinePlugin({
          __VERSION__: JSON.stringify('dev')
        })
      ]
  },
  devServer: {
    watchOptions: {
      poll: 1000
    }
  }
}
