package wix

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

func GetWixOrders() (WixOrdersResponse, error) {
	var orders WixOrdersResponse
	jsonFile, err := os.Open("raworder.json")
	if err != nil {
		return WixOrdersResponse{}, err
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	err = json.Unmarshal(byteValue, &orders)
	if err != nil {
		return WixOrdersResponse{}, err
	}
	return orders, nil
}
