package credentials

import (
	"encoding/json"
	"io"
)

type IBMCloud struct{}

var config *CloudLogin
var iHttp IBMHttp

func (i *IBMCloud) init() {
	env := Enviroment{}
	config = env.GetConfig()
	iHttp = IBMHttp{}
}

func (i *IBMCloud) getResult() []byte {
	resp := iHttp.GetResponse()
	result, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	return result
}

func (i *IBMCloud) parseToken(result []byte) IBMToken {
	var token IBMToken
	err := json.Unmarshal(result, &token)
	if err != nil {
		panic(err)
	}
	return token
}

func (i *IBMCloud) GetToken() IBMToken {
	i.init()
	result := i.getResult()
	return i.parseToken(result)
}
