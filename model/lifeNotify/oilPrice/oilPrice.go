package oilPrice

import (
	"encoding/json"
	"lkserver/model/httpHandler"
)

var (
	oilInfo Message
)

const (
	//Key
	apiKey string = "ae76ca2c0a933113be2bd17ab9b221f6"
	//天行数据-实时油价接口
	oilBaseUrl string = "https://api.tianapi.com/oilprice/index?"
	//省份
	province string = "浙江"
)

type OilData struct {
	Prov string `json:"prov"`
	P0   string `json:"p0"`
	P89  string `json:"p89"`
	P92  string `json:"p92"`
	P95  string `json:"p95"`
	P98  string `json:"p98"`
	Time string `json:"time"`
}

type Message struct {
	Code     int       `json:"code"`
	Msg      string    `json:"msg"`
	Newslist []OilData `json:"newslist"`
}

//获取油价
func GetOilInfoFromTianxing() {
	url := oilBaseUrl + "key=" + apiKey + "&prov=" + province
	body, _ := httpHandler.SendHttpGetMessage(url)

	json.Unmarshal(body, &oilInfo)
}

//获取92号油价
func GetOilPrice92() string {
	return oilInfo.Newslist[0].P92
}

//获取95号油价
func GetOilPrice95() string {
	return oilInfo.Newslist[0].P95
}

//组装成股票通知消息
func GetOilInfoNotify() string {
	notidfy := province + "油价：" + "92号(" + GetOilPrice92() + ") 95号(" + GetOilPrice95() + ")"
	return notidfy
}

//直接获取一次股票信息
func GetStockPriceOnce() string {
	url := oilBaseUrl + "key=" + apiKey + "&prov=" + province
	body, _ := httpHandler.SendHttpGetMessage(url)

	json.Unmarshal(body, &oilInfo)

	return GetOilInfoNotify()
}
