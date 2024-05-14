package govite

import (
	"encoding/json"
	"fmt"
	viteManifest "github.com/aquaswim/govite/vite_manifest"
	"io/fs"
	"os"
	"path"
)

type ViteAdapter struct {
	cfg           *Config
	isDevelopment bool
	manifest      viteManifest.T
	fs            fs.FS
}

func New(cfg *Config) *ViteAdapter {
	return NewWithFS(cfg, defaultFs)
}

func NewWithFS(cfg *Config, fs fs.FS) *ViteAdapter {
	adapter := ViteAdapter{
		cfg: cfg,
		fs:  fs,
	}
	// validate vite output dir
	mustValidateOutputDir(cfg.ViteOutputPath)
	// determine mode
	devManifest, err := fs.Open(path.Join(cfg.ViteOutputPath, "manifest.dev.json"))
	if err == nil {
		defer devManifest.Close()
		fmt.Println("Vite Dev Mode")
		adapter.isDevelopment = true
		// get the manifest
		adapter.manifest = &viteManifest.Dev{}
		err := json.NewDecoder(devManifest).Decode(adapter.manifest)
		if err != nil {
			panic(err)
		}
		return &adapter
	}

	prodManifest, err := fs.Open(path.Join(cfg.ViteOutputPath, ".vite", "manifest.json"))
	if err != nil {
		panic("Manifest file not found!")
	}
	defer prodManifest.Close()
	fmt.Println("Vite Prod Mode")
	// get the manifest
	adapter.manifest = &viteManifest.Prod{}
	err = json.NewDecoder(prodManifest).Decode(adapter.manifest)
	if err != nil {
		panic(err)
	}

	return &adapter
}

func mustValidateOutputDir(path string) {
	stat, err := os.Stat(path)
	if err != nil {
		panic(err)
	}
	if !stat.IsDir() {
		panic(fmt.Sprintf("%s is not valid directory", path))
	}
}

func (v *ViteAdapter) GetBuilder() (VitePageAssetsBuilder, error) {
	if v.isDevelopment {
		return newVitePageAssetDev(v.manifest.(*viteManifest.Dev)), nil
	}

	return newVitePageAssetProd(v.manifest), nil
}

func (v *ViteAdapter) MustGetBuilder() VitePageAssetsBuilder {
	builder, err := v.GetBuilder()
	if err != nil {
		panic(err)
	}
	return builder
}
