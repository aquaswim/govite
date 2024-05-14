package vite_manifest

import "log"

type Prod map[string]Chunk

func (p Prod) GetViteServerUrl() (string, error) {
	log.Panicln("Prod manifest don't have server url")
	return "", nil
}

func (p Prod) Resolve(name string) *Chunk {
	val, ok := p[name]
	if ok {
		return &val
	}
	return nil
}
