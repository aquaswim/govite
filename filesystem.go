package govite

import (
	"io/fs"
	"os"
)

type osFS struct{}

func (o osFS) Open(name string) (fs.File, error) {
	return os.Open(name)
}

var defaultFs osFS = osFS{}
