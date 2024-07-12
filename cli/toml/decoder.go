package toml

import "github.com/BurntSushi/toml"

func Decode(path string, v any) (toml.MetaData, error) {
	return toml.DecodeFile(path, v)
}
