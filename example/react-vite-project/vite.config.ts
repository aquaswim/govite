import {defineConfig} from 'vite'
import react from '@vitejs/plugin-react'
import devManifest from 'vite-plugin-dev-manifest'

// https://vite.dev/config/
export default defineConfig({
    plugins: [react(), devManifest()],
    build: {
        manifest: true,
        rollupOptions: {
            input: './src/main.tsx',
        },
    },
})
