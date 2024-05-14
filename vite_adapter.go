package govite

import (
	"encoding/json"
	viteManifest "github.com/aquaswim/govite/vite_manifest"
	"io/fs"
	"log"
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
	mustValidateOutputDir(fs, path.Clean(cfg.ViteOutputPath))
	// determine mode
	devManifest, err := fs.Open(path.Join(cfg.ViteOutputPath, "manifest.dev.json"))
	if err == nil {
		defer devManifest.Close()
		log.Println("Vite Dev Mode")
		adapter.isDevelopment = true
		// get the manifest
		adapter.manifest = &viteManifest.Dev{}
		err := json.NewDecoder(devManifest).Decode(adapter.manifest)
		if err != nil {
			log.Panicf("error reading vite dev manifest, error: %s", err)
		}
		return &adapter
	}

	prodManifest, err := fs.Open(path.Join(cfg.ViteOutputPath, ".vite", "manifest.json"))
	if err != nil {
		log.Panicf("Manifest file not found!")
	}
	defer prodManifest.Close()
	log.Println("Vite Prod Mode")
	// get the manifest
	adapter.manifest = &viteManifest.Prod{}
	err = json.NewDecoder(prodManifest).Decode(adapter.manifest)
	if err != nil {
		log.Panicf("error reading vite prod manifest: %s", err)
	}

	return &adapter
}

func mustValidateOutputDir(fs fs.FS, path string) {
	file, err := fs.Open(path)
	if err != nil {
		log.Panicln(err)
	}
	defer file.Close()
	stat, err := file.Stat()
	if err != nil {
		log.Panicln(err)
	}
	if !stat.IsDir() {
		log.Panicf("%s is not valid directory", path)
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
		log.Panic(err)
	}
	return builder
}
