package main

import (
	"context"
	"io"

	http_utils "jerry.io/core/http"
	http_types "jerry.io/core/http/types"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) DoRequest(url string, method string) http_types.Response {
	res := http_types.Response{}
	response, err := http_utils.MakeRequest(http_types.Request{
		Url:    url,
		Method: method,
	})
	if err != nil {
		// TODO: For now
		return res
	}
	res.Status = response.Status
	res.StatusCode = response.StatusCode
	defer response.Body.Close()
	bodyData, err := io.ReadAll(response.Body)
	if err != nil {
		return res
	}
	res.Body = string(bodyData)
	return res
}
