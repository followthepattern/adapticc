const path = require('path');
const HtmlWebpackPlugin = require('html-webpack-plugin');
const { TailwindCSSWebpackPlugin } = require('tailwindcss-webpack-plugin')
const Dotenv = require('dotenv-webpack');

module.exports = (env, argv) => {
  console.info("env", env);
  console.info("argv", argv);
  return {
    entry: './app/index.tsx',
    output: {
      path: path.resolve(__dirname, 'dist'),
      filename: 'bundle.js',
      publicPath: "/",
    },
    module: {
      rules: [
        {
          test: /\.(ts|tsx)$/,
          use: {
            loader: 'ts-loader',
          },
        },
        {
          test: /\.css$/i,
          use: ["style-loader", "css-loader"],
        },
      ],
    },
    devServer: {
      static: {
        directory: path.join(__dirname, 'public'),
      },
      historyApiFallback: true,
      compress: true,
      port: 3000,
    },
    plugins: [
      new HtmlWebpackPlugin({
        template: './public/index.html',
      }),
      new TailwindCSSWebpackPlugin({
        devtools: {
          host: '0.0.0.0'
        }
      }),
      new Dotenv({
        path: `.env.${argv.mode}`
      }),
    ],
    resolve: {
      alias: {
        "react-router-dom": path.resolve(
          __dirname,
          "node_modules/react-router-dom/dist/index"
        ),
        "@": path.resolve(
          __dirname,
          "./"),
      },
      extensions: ['.ts', '.tsx', '.js', '.jsx'],
    },
  };
};
