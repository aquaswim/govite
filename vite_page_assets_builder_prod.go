package govite

import (
	"fmt"
	viteManifest "github.com/aquaswim/govite/vite_manifest"
)

type vitePageAssetsProd struct {
	manifest viteManifest.T
	baseUrl  string

	Styles        []string
	ScriptModule  []string
	ModulePreload []string
}

func newVitePageAssetProd(manifest viteManifest.T) VitePageAssetsBuilder {
	return &vitePageAssetsProd{
		manifest:      manifest,
		Styles:        []string{},
		ScriptModule:  []string{},
		ModulePreload: []string{},
	}
}

func (a *vitePageAssetsProd) AddAsset(assetFile string) VitePageAssetsBuilder {
	a.addAsset(assetFile, false)
	return a
}

func (a *vitePageAssetsProd) addAsset(assetFile string, imported bool) {
	chunk := a.manifest.Resolve(assetFile)
	if chunk != nil {
		// append css
		if len(chunk.Css) > 0 {
			for _, css := range chunk.Css {
				a.Styles = appendIfNotExists(a.Styles, css)
			}
		}
		// append file
		if chunk.File != "" {
			if imported {
				a.ModulePreload = appendIfNotExists(a.ModulePreload, chunk.File)
			} else {
				a.ScriptModule = appendIfNotExists(a.ScriptModule, chunk.File)
			}
		}

		// imports
		for _, _import := range chunk.Imports {
			// todo: i think infinite loop can happen
			a.addAsset(_import, true)
		}
	}
}

func (a *vitePageAssetsProd) CreateStyleTags() string {
	var out string
	for _, style := range a.Styles {
		out += fmt.Sprintf(`<link rel="stylesheet" href="%s/%s" />`, a.baseUrl, style)
	}
	return out
}

func (a *vitePageAssetsProd) CreateScriptTags() string {
	var out string
	for _, script := range a.ScriptModule {
		out += fmt.Sprintf(`<script type="module" src="%s/%s"></script>`, a.baseUrl, script)
	}
	return out
}

func (a *vitePageAssetsProd) CreatePreloadTags() string {
	var out string
	for _, preload := range a.ModulePreload {
		out += fmt.Sprintf(`<link rel="modulepreload" href="%s%s" />`, a.baseUrl, preload)
	}
	return out
}
