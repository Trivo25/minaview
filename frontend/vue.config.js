module.exports = {
  devServer: {
    proxy: {
      '/': {
        target: 'http://localhost:8000',
        ws: true,
        changeOrigin: true,
      },
    },
  },
  publicPath: './',
}
