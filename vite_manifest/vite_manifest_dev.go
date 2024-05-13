package vite_manifest

type Dev struct {
	Url    string            `json:"url"`
	Inputs map[string]string `json:"inputs"`
}

func (d Dev) GetViteServerUrl() (string, error) {
	return d.Url, nil
}

func (d Dev) Resolve(_ string) *Chunk {
	panic("dev manifest cannot be resolved")
}
