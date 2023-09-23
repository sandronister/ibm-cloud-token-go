package main

import (
	"encoding/json"
	"os"

	"lab.ibm.cloud/credentials"
)

func main() {
	iCloud := credentials.IBMCloud{}
	res := iCloud.GetToken()
	json.NewEncoder(os.Stdout).Encode(res)
}
