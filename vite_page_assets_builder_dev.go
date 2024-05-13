package govite

import (
	"fmt"
	viteManifest "github.com/aquaswim/govite/vite_manifest"
)

type vitePageAssetsDev struct {
	*viteManifest.Dev
	input []string
}

func newVitePageAssetDev(manifest *viteManifest.Dev) VitePageAssetsBuilder {
	return &vitePageAssetsDev{
		Dev:   manifest,
		input: []string{},
	}
}

func (v *vitePageAssetsDev) AddAsset(assetFile string) VitePageAssetsBuilder {
	for name, url := range v.Inputs {
		if name == assetFile || url == assetFile {
			v.input = appendIfNotExists(v.input, assetFile)
			break
		}
	}
	return v
}

func (v *vitePageAssetsDev) CreateStyleTags() string {
	return ""
}

func (v *vitePageAssetsDev) CreateScriptTags() string {
	out := "<!-- Vite Development -->"
	out += fmt.Sprintf(`<script type="module" src="%s@vite/client"></script>`, v.Url)
	for _, script := range v.input {
		out += fmt.Sprintf(`<script type="module" src="%s%s"></script>`, v.Url, script)
	}
	return out
}

func (v *vitePageAssetsDev) CreatePreloadTags() string {
	return ""
}
