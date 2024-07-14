package http_types

type TOMLFile struct {
	Request Request `toml:"request"`
}

type Request struct {
	Url     string   `toml:"url"`
	Method  string   `toml:"method"`
	Headers []Header `toml:"headers"`
	Body    string   `toml:"body"`
}
type Header struct {
	Name  string
	Value string
}

type Response struct {
	Status     string `json:"status"`
	StatusCode int    `json:"statusCode"`
	Body       string `json:"body"`
}
