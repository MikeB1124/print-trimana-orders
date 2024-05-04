package wix

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/MikeB1124/print-trimana-orders/configuration"
	"github.com/MikeB1124/print-trimana-orders/logger"
)

func GetWixOrders() (WixOrdersResponse, error) {
	logger.InfoLogger.Println("Creating new GET request to pull wix orders")
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/orders?status=NEW", configuration.Config.WixConfig.Url), nil)
	if err != nil {
		return WixOrdersResponse{}, err
	}
	logger.InfoLogger.Println("Set http request headers for wix api call")
	req.Header.Set("Authorization", configuration.Config.WixConfig.Auth)
	req.Header.Set("wix-account-id", configuration.Config.WixConfig.AccountID)
	req.Header.Set("wix-site-id", configuration.Config.WixConfig.SiteID)

	//Execute http request
	logger.InfoLogger.Println("Execute http GET request for listing wix orders")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return WixOrdersResponse{}, err
	}

	//Read http response body bytes
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return WixOrdersResponse{}, err
	}

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		//Unmarshal response to struct to return to client
		logger.InfoLogger.Println("Unmarshal response for wix orders if 200 status code")
		var resData WixOrdersResponse
		errUnmarshal := json.Unmarshal(bodyBytes, &resData)
		if errUnmarshal != nil {
			return WixOrdersResponse{}, err
		}
		return resData, nil
	} else {
		return WixOrdersResponse{}, fmt.Errorf("Error: http status code was %d", resp.StatusCode)
	}
}
