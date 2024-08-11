package http_types

type TOMLFile struct {
	Request Request `toml:"request"`
}
type Response struct {
	Status     string `json:"status"`
	StatusCode int    `json:"statusCode"`
	Body       string `json:"body"`
}
