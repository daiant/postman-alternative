package toml

type TOMLFile struct {
	Request Request `toml:"request"`
}
type Request struct {
	Url     string   `toml:"url"`
	Method  string   `toml:"method"`
	Headers []Header `toml:"headers"`
}
type Header struct {
	Name  string
	Value string
}
