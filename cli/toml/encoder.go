package toml

import (
	"io"

	"github.com/BurntSushi/toml"
)

func Encode(f io.Writer, data TOMLFile) error {
	return toml.NewEncoder(f).Encode(data)
}
