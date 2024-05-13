# go-vite

Vite Adapter in go

# How to use

1. install this plugins: [vite-plugin-dev-manifest](https://github.com/owlsdepartment/vite-plugin-dev-manifest) in package.json
2. add the vite-plugin-dev-manifest to vite.config and enable manifest generation.
    ```js
   import {defineConfig} from "vite";
   import devManifest from 'vite-plugin-dev-manifest';

    export default defineConfig({
        plugins: [
            //.. add your other plugins here ..
            devManifest(),
        ],
        build: {
        // generate .vite/manifest.json in outDir
        manifest: true,
        rollupOptions: {
                // your scripts
                input: './main.js',
            },
        },
    })
    ```
3. to implement the server you can see the example at `example/simple-vite-project/server.go`