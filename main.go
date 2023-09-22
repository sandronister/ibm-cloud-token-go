package main

import (
	"io"
	"net/http"
	"net/url"
	"strings"

	"lab.bia.env2/credentials"
)

var config credentials.IBMConfig

func getForm() *strings.Reader {
	form := url.Values{}
	form.Add("grant_type", config.GranType)
	form.Add("apikey", config.ApiKey)
	return strings.NewReader(form.Encode())
}

func getRequest() *http.Request {
	config = credentials.LoadConfig()
	form := getForm()
	req, err := http.NewRequest("POST", config.Url, form)
	if err != nil {
		panic(err)
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	return req
}

func getResult(req *http.Request) []byte {
	c := http.Client{}
	resp, err := c.Do(req)
	if err != nil {
		panic(err)
	}
	result, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	return result
}

func main() {
	req := getRequest()
	result := getResult(req)
	println(string(result))
}
