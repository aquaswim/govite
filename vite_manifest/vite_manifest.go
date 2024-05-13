package vite_manifest

type T interface {
	GetViteServerUrl() (string, error)
	Resolve(name string) *Chunk
}

type Chunk struct {
	File           string   `json:"file"`
	Src            string   `json:"src"`
	IsEntry        bool     `json:"isEntry"`
	DynamicImports []string `json:"dynamicImports"`
	Css            []string `json:"css"`
	Assets         []string `json:"assets"`
	Imports        []string `json:"imports"`
}
