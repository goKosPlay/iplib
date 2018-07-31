package iplib

type IpDetail struct {
	Ip        string `json:"ip"`
	City      string `json:"city"`
	Region    interface{} `json:"region"`
	Country   string `json:"country"`
	Postal    interface{} `json:"postal"`
	Latitude  float32 `json:"latitude"`
	Longitude float32 `json:"longitude"`
	Timezone  string `json:"timezone"`
}
