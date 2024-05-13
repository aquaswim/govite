package govite

type VitePageAssetsBuilder interface {
	CreateStyleTags() string
	CreateScriptTags() string
	CreatePreloadTags() string

	AddAsset(assetFile string) VitePageAssetsBuilder
}
