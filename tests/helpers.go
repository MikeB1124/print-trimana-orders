package tests

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/MikeB1124/print-trimana-orders/wix"
)

func GetJsonTestOrders() (wix.WixOrdersResponse, error) {
	var orders wix.WixOrdersResponse
	jsonFile, err := os.Open("raworder.json")
	if err != nil {
		return wix.WixOrdersResponse{}, err
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	err = json.Unmarshal(byteValue, &orders)
	if err != nil {
		return wix.WixOrdersResponse{}, err
	}
	return orders, nil
}
