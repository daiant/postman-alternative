package http

import (
	"bytes"
	"cli/toml"
	"fmt"
	"io"
	"net/http"
	"os"
	filepathUtil "path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

func init() {
	HttpCmd.Flags().StringVarP(&url, "url", "u", "", "URL to fetch data from")
	HttpCmd.Flags().StringVarP(&method, "method", "X", "GET", "HTTP Method")
	HttpCmd.Flags().StringArrayVarP(&headers, "header", "H", []string{}, "Header present in the payload")
	HttpCmd.Flags().StringVar(&body, "data", "", "POST Body")
	HttpCmd.Flags().StringVarP(&filepath, "file", "f", "", "Location of TOML Request File")
	HttpCmd.Flags().StringVar(&save, "save", "", "Saves request to configuration file")
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
	save     string
	HttpCmd  = &cobra.Command{
		Use:   "http",
		Short: "Makes an HTTP request",
		Run: func(cmd *cobra.Command, args []string) {
			if len(filepath) > 0 {
				loadFromFile()
			}
			printResult(makeRequest(method))
			if len(save) > 0 {
				saveToFile()
			}
		},
	}
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}
func getExecutableDir() string {
	ex, err := os.Executable()
	check(err)
	return filepathUtil.Dir(ex)
}
func saveToFile() {
	var tomlData toml.TOMLFile
	tomlData.Request.Url = url
	tomlData.Request.Method = method
	for _, v := range headers {
		h := strings.Split(v, ":")
		tomlData.Request.Headers = append(tomlData.Request.Headers, toml.Header{Name: strings.TrimSpace(h[0]), Value: strings.TrimSpace(h[1])})
	}

	file, err := os.Create(getExecutableDir() + "/storage/" + save + ".toml")
	check(err)
	defer file.Close()
	toml.Encode(file, tomlData)
	check(err)

	fmt.Printf("Saved request in %s\n", getExecutableDir()+"/storage/"+save+".toml")
}

func loadFromFile() {
	wd, err := os.Getwd()
	check(err)

	var tomlData toml.TOMLFile
	toml.Decode(wd+"/"+filepath, &tomlData)
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
