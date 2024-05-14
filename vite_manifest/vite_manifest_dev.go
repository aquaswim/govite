package vite_manifest

import "log"

type Dev struct {
	Url    string            `json:"url"`
	Inputs map[string]string `json:"inputs"`
}

func (d Dev) GetViteServerUrl() (string, error) {
	return d.Url, nil
}

func (d Dev) Resolve(_ string) *Chunk {
	log.Panicf("dev manifest cannot be resolved")
	return nil
}
