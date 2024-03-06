import { defineConfig } from 'vite'
import react from '@vitejs/plugin-react'

export default defineConfig({
  
  plugins: [react()],
  server: {
    proxy: {
      '/api': {
        target: 'http://localhost:8000', // Update with the URL of your Go backend
        changeOrigin: true,
        secure: false,
      },
    },
  },
})
