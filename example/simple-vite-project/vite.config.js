import {defineConfig} from "vite";
import devManifest from 'vite-plugin-dev-manifest';

// vite.config.js
export default defineConfig({
    plugins: [
        devManifest(),
    ],
    build: {
        // generate .vite/manifest.json in outDir
        manifest: true,
        rollupOptions: {
            // overwrite default .html entry
            input: './main.js',
        },
    },
})