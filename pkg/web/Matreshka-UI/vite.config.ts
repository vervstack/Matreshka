import { fileURLToPath, URL } from 'node:url'

import {defineConfig, loadEnv} from 'vite'
import vue from '@vitejs/plugin-vue'

// @ts-expect-error
export default ({mode}) => {
  process.env = {...process.env, ...loadEnv(mode, process.cwd())};

  return defineConfig({
    base: '/',
    plugins: [vue()],
    resolve: {
      alias: {
        '@': fileURLToPath(new URL('./src', import.meta.url)),
        // '@vervstack/matreshka': fileURLToPath(
        //     new URL('../@vervstack/matreshka/dist/index.js', import.meta.url)),
      },
      dedupe: ['@vervstack/matreshka'],
    },
    optimizeDeps: {
      include: ['@vervstack/matreshka']
    },
    build: {
      rollupOptions: {
        output: {
          entryFileNames: '[name].js',
          chunkFileNames: '[name].js',
          assetFileNames: '[name].[ext]'
        }
      }
    },
  });
}
