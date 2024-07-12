package http

import (
	"bytes"
	"cli/toml"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

func init() {
	HttpCmd.Flags().StringVarP(&url, "url", "u", "", "URL to fetch data from")
	HttpCmd.Flags().StringVarP(&method, "method", "X", "GET", "HTTP Method")
	HttpCmd.Flags().StringArrayVarP(&headers, "header", "H", []string{}, "Header present in the payload")
	HttpCmd.Flags().StringVar(&body, "data", "", "POST Body")
	HttpCmd.Flags().StringVarP(&filepath, "file", "f", "", "Location of TOML Request File")

	HttpCmd.MarkFlagsOneRequired("url", "file")
	client = http.Client{}
}

var (
	client   http.Client
	body     string
	url      string
	filepath string
	method   string
	headers  []string
	HttpCmd  = &cobra.Command{
		Use:   "http",
		Short: "Makes an HTTP request",
		Run: func(cmd *cobra.Command, args []string) {
			if len(url) <= 0 {
				loadFromFile()
			}
			printResult(makeRequest(method))
		},
	}
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}
func loadFromFile() {
	wd, err := os.Getwd()
	check(err)

	var tomlData toml.TOMLFile
	toml.Decode(wd+filepath, &tomlData)
	check(err)

	url = tomlData.Request.Url
	method = tomlData.Request.Method

	for _, h := range tomlData.Request.Headers {
		headers = append(headers, h.Name+":"+h.Value)
	}
}

func makeRequest(method string) (*http.Response, error) {
	request, err := http.NewRequest(method, url, bytes.NewReader([]byte(body)))
	if err != nil {
		return nil, err
	}
	setHeaders(request)
	return client.Do(request)
}

func setHeaders(request *http.Request) {
	for _, header := range headers {
		parts := strings.Split(header, ":")
		if len(parts) == 2 {
			request.Header.Add(strings.TrimSpace(parts[0]), strings.TrimSpace(parts[1]))
		}
	}
}

func printResult(response *http.Response, err error) {
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Status: %d\n\n", response.StatusCode)
		for key, values := range response.Header {
			fmt.Printf("%s:\t%s\n", key, strings.Join(values, ", "))
		}
		defer response.Body.Close()
		fmt.Println("Response:")
		bodyBytes, err := io.ReadAll(response.Body)
		if err != nil {
			fmt.Println(err)
		} else {
			bodyString := string(bodyBytes)
			fmt.Println(bodyString)
		}
	}
}
