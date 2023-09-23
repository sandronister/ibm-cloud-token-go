package credentials

import (
	"net/http"
	"net/url"
	"strings"
)

type IBMHttp struct{}

func (i *IBMHttp) getParams() *strings.Reader {
	params := url.Values{}
	params.Add("grant_type", config.GranType)
	params.Add("apikey", config.ApiKey)
	return strings.NewReader(params.Encode())
}

func (i *IBMHttp) getRequest() *http.Request {
	params := i.getParams()
	req, err := http.NewRequest("POST", config.Url, params)
	if err != nil {
		panic(err)
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	return req
}

func (i *IBMHttp) GetResponse() *http.Response {
	req := i.getRequest()
	c := http.Client{}
	resp, err := c.Do(req)
	if err != nil {
		panic(err)
	}
	return resp
}
