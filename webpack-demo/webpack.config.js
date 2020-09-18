const { CleanWebpackPlugin } = require('clean-webpack-plugin');
const HtmlWebpackPlugin = require('html-webpack-plugin');
const path = require('path');

module.exports = {
  mode: process.env.NODE_ENV || 'development',
  entry: path.resolve(__dirname, 'src/index.js'),
  devtool: 'inline-source-map',
  devServer: {
    contentBase: path.resolve(__dirname, 'dist'),
  },
  plugins: [
    new CleanWebpackPlugin({
      cleanStaleWebpackAssets: false,
    }),
    new HtmlWebpackPlugin({
      title: 'Caching',
    }),
  ],
  output: {
    filename: '[name].[contenthash].js',
    path: path.resolve(__dirname, 'dist'),
    publicPath: '/',
  },
  optimization: {
    moduleIds: 'hashed',
    runtimeChunk: 'single',
    splitChunks: {
      cacheGroups: {
        vendor: {
          test: /[\\/]node_modules[\\/]/,
          name: 'vendors',
          chunks: 'all',
        },
      },
    },
  },
  module: {
    rules: [{
      test: /\.css$/,
      use: [
        'style-loader',
        'css-loader',
      ],
    }, {
      test: /\.(png|svg|jpg|gif)$/,
      loader: 'file-loader',
    }, {
      test: /\.(woff|woff2|eot|ttf|otf)$/,
      loader: 'file-loader',
    }, {
      test: /\.(csv|tsv)$/,
      loader: 'csv-loader',
    }, {
      test: /\.xml$/,
      loader: 'xml-loader',
    }],
  },
};
