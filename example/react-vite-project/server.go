package main

import (
	"fmt"
	"github.com/aquaswim/govite"
	"log"
	"net/http"
)

func main() {
	vite := govite.New(&govite.Config{
		ViteOutputPath: "./dist",
		AssetBaseUrl:   "/assets",
		IsReact:        true,
	})

	mux := http.NewServeMux()

	mux.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		b := vite.MustGetBuilder().AddAsset("src/main.tsx")

		w.Write([]byte(fmt.Sprintf(`
			<!doctype html>
			<html lang="en">
				<head>
					<meta charset="UTF-8" />
					<link rel="icon" type="image/svg+xml" href="/vite.svg" />
					<meta name="viewport" content="width=device-width, initial-scale=1.0" />
					<title>Vite + React + TS</title>
					%s
					%s
				</head>
				<body>
					<div id="root"></div>
					%s
				</body>
			</html>
	`, b.CreateStyleTags(), b.CreatePreloadTags(), b.CreateScriptTags())))
	})

	// serve the dist folder
	mux.Handle("/", vite.FileServer())

	log.Println("Listening: http://localhost:3000/home")
	err := http.ListenAndServe(":3000", mux)
	if err != nil {
		panic(err)
	}
}
