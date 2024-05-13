package vite_manifest

type Prod map[string]Chunk

func (p Prod) GetViteServerUrl() (string, error) {
	panic("Prod manifest don't have server url")
}

func (p Prod) Resolve(name string) *Chunk {
	val, ok := p[name]
	if ok {
		return &val
	}
	return nil
}
