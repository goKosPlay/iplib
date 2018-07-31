package iplib

import (
	"net/http"
	"io/ioutil"
	"encoding/json"
	"errors"
	"fmt"
)

//ip detail

type IpMapData map[string]interface{}
type IpMod struct {
	IpAddress NowIP
	IpInfo    IpDetail
}

func (self *IpMod) GetIp() string {
	data, err := getResponseData("https://api.ipify.org?format=json", self.IpAddress)
	if err != nil {
		panic(err)
	}
	ip := fmt.Sprintf("%v", data.(IpMapData)["ip"])
	return ip
}

func (self *IpMod) GetIpDetail() map[string]string {
	data, err := getResponseData("https://ipapi.co/"+self.GetIp()+"/json", self.IpInfo)
	if err != nil {
		panic(err)
	}
	ipDescription := make(map[string]string)
	for k, v := range data.(IpMapData) {
		ipDescription[k] = fmt.Sprintf("%v", v)
	}
	return ipDescription
}

func (self *IpMod) GetSelectIpDetail(ip string) map[string]string {
	data, err := getResponseData("https://ipapi.co/"+ip+"/json", self.IpInfo)
	if err != nil {
		panic(err)
	}
	ipDescription := make(map[string]string)
	for k, v := range data.(IpMapData) {
		ipDescription[k] = fmt.Sprintf("%v", v)
	}
	return ipDescription
}

func NewIpMod() *IpMod {
	return &IpMod{}
}

func getResponseData(url string, jsonStruct interface{}) (interface{}, error) {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	if err := json.Unmarshal([]byte(body), &jsonStruct); err == nil {
		return jsonStruct, nil
	}
	return nil, errors.New("not found the object")
}