package iplib

import (
	"net/http"
	"io/ioutil"
	"encoding/json"
	"errors"
)

//https://ipapi.co/0.0.0.0/json
//{
//"ip": "0.0.0.0",
//"city": "",
//"region": null,
//"country": "",
//"postal": null,
//"latitude": 0,
//"longitude": 0,
//"timezone": "Asia/Taipei"
//}
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

//https://api.ipify.org?format=json
//{"ip":"0.0.0.0"}
type MyIp struct {
	Ip string `json:"ip"`
}

//取得真實IP
func GetRealIpAddress() IpInfo {
	var myIp MyIp
	var ipInfo IpInfo
	getMyInfo("https://api.ipify.org?format=json", &myIp)
	getMyInfo("https://ipapi.co/"+myIp.Ip+"/json", &ipInfo)
	return ipInfo
}

func getMyInfo(url string, myInterface interface{}) (interface{}, error) {
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
