import { defineConfig } from 'vite'
import react from '@vitejs/plugin-react'
import { TanStackRouterVite } from '@tanstack/router-plugin/vite'

// https://vitejs.dev/config/
export default defineConfig({
  base: "/",
  plugins: [
    react(),
    TanStackRouterVite(),
  ],
  preview: {
    port: 8000,
    strictPort: true,
  },
  server: {
    port: 8000,
    strictPort: true,
    host: true,
  },
})
