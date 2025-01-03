package govite

import (
	"fmt"
	viteManifest "github.com/aquaswim/govite/vite_manifest"
)

type vitePageAssetsDev struct {
	*viteManifest.Dev
	input   []string
	isReact bool
}

func newVitePageAssetDev(manifest *viteManifest.Dev, isReact bool) VitePageAssetsBuilder {
	return &vitePageAssetsDev{
		Dev:     manifest,
		input:   []string{},
		isReact: isReact,
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
	if v.isReact {
		// add react render hook
		out += `<script type="module">
  import RefreshRuntime from '` + v.Url + `@react-refresh'
  RefreshRuntime.injectIntoGlobalHook(window)
  window.$RefreshReg$ = () => {}
  window.$RefreshSig$ = () => (type) => type
  window.__vite_plugin_react_preamble_installed__ = true
</script>`
	}
	out += fmt.Sprintf(`<script type="module" src="%s@vite/client"></script>`, v.Url)
	for _, script := range v.input {
		out += fmt.Sprintf(`<script type="module" src="%s%s"></script>`, v.Url, script)
	}
	return out
}

func (v *vitePageAssetsDev) CreatePreloadTags() string {
	return ""
}
