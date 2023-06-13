import { defineConfig } from 'vite'
import react from '@vitejs/plugin-react'
import * as path from "path";
import { NodeGlobalsPolyfillPlugin } from '@esbuild-plugins/node-globals-polyfill'
// import commonjsPlugin from '@chialab/esbuild-plugin-commonjs';

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [
    react(),
    NodeGlobalsPolyfillPlugin({
      buffer: true
    })
  ],
  define: {
    'process.env': {},
    global: 'globalThis'
  },
  resolve: { alias: { web3: path.resolve(__dirname, '../../node_modules/web3/dist/web3.min.js') }, }
})
