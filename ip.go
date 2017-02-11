package iplib

import (
	"net/http"
	"io/ioutil"
	"encoding/json"
	"errors"
)

//ip detail
type IpInfo struct {
	Ip        string `json:"ip"`
	City      string `json:"city"`
	Region    interface{} `json:"region"`
	Country   string `json:"country"`
	Postal    interface{} `json:"postal"`
	Latitude  float32 `json:"latitude"`
	Longitude float32 `json:"longitude"`
	Timezone  string `json:"timezone"`
}

type MyIp struct {
	Ip string `json:"ip"`
}

//取得真實IP
func GetRealIpAddress() IpInfo {
	var myIp MyIp
	var ipInfo IpInfo
	getResponseData("https://api.ipify.org?format=json", &myIp)
	getResponseData("https://ipapi.co/"+myIp.Ip+"/json", &ipInfo)
	return ipInfo
}

//取得IP資訊
func GetIpInformation(strIp string) IpInfo {
	var ipinfo IpInfo
	getResponseData("https://ipapi.co/"+strIp+"/json", &ipinfo)
	return ipinfo
}

func getResponseData(url string, myInterface interface{}) (interface{}, error) {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	if err := json.Unmarshal([]byte(body), &myInterface); err == nil {
		return myInterface, nil
	}
	return nil, errors.New("not found the object")
}

func GetVersion() string {
	return "0.1"
}