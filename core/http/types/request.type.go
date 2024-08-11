package http_types

type Request struct {
	Url     string     `json:"url"`
	Method  string     `json:"method"`
	Body    string     `json:"body"`
	Params  [][]string `json:"params"`
	Headers [][]string `json:"headers"`
}
