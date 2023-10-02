package main

import (
	"encoding/json"
	"os"

	"github.com/gofor-little/env"
	"github.com/sandronister/ibm-cloud-token-go/adapter/ibm/di"
)

func init() {
	err := env.Load(".env")
	if err != nil {
		panic(err)
	}
}

func main() {
	ibmConnect := di.ConfigIBMConect()
	token, err := ibmConnect.GetToken()
	if err != nil {
		panic(err)
	}
	json.NewEncoder(os.Stdout).Encode(token)
}
