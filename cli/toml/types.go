package toml

type TOMLFile struct {
	Request struct {
		Url     string   `toml:"url"`
		Method  string   `toml:"method"`
		Headers []header `toml:"headers"`
	}
}
type header struct {
	Name  string
	Value string
}
