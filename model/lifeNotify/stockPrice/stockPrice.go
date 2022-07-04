package stockPrice

import (
	"encoding/json"
	"lkserver/model/component"
	"lkserver/model/httpHandler"
	"strconv"
)

var (
	stockInfo Message
)

const (
	//股票代码
	stockCode string = "SZ000925"
	//雪球股票接口
	stockBaseUrl string = "https://stock.xueqiu.com/v5/stock/realtime/quotec.json?"
)

type StockData struct {
	Symbol             string      `json:"symbol"`
	Current            float64     `json:"current"`
	Percent            float64     `json:"percent"`
	Chg                float64     `json:"chg"`
	Timestamp          int64       `json:"timestamp"`
	Volume             int         `json:"volume"`
	Amount             float64     `json:"amount"`
	MarketCapital      int64       `json:"market_capital"`
	FloatMarketCapital int64       `json:"float_market_capital"`
	TurnoverRate       float64     `json:"turnover_rate"`
	Amplitude          float64     `json:"amplitude"`
	Open               float64     `json:"open"`
	LastClose          float64     `json:"last_close"`
	High               float64     `json:"high"`
	Low                float64     `json:"low"`
	AvgPrice           float64     `json:"avg_price"`
	TradeVolume        interface{} `json:"trade_volume"`
	Side               int         `json:"side"`
	IsTrade            bool        `json:"is_trade"`
	Level              int         `json:"level"`
	TradeSession       interface{} `json:"trade_session"`
	TradeType          interface{} `json:"trade_type"`
	CurrentYearPercent float64     `json:"current_year_percent"`
	TradeUniqueID      interface{} `json:"trade_unique_id"`
	Type               int         `json:"type"`
	BidApplSeqNum      interface{} `json:"bid_appl_seq_num"`
	OfferApplSeqNum    interface{} `json:"offer_appl_seq_num"`
	VolumeExt          interface{} `json:"volume_ext"`
	TradedAmountExt    interface{} `json:"traded_amount_ext"`
	TradeTypeV2        interface{} `json:"trade_type_v2"`
}

type Message struct {
	Data             []StockData `json:"data"`
	ErrorCode        int         `json:"error_code"`
	ErrorDescription interface{} `json:"error_description"`
}

//获取股票价格
func GetStockInfoFromXueqiu() {
	url := stockBaseUrl + "symbol=" + stockCode
	body, _ := httpHandler.SendHttpGetMessage(url)

	json.Unmarshal(body, &stockInfo)
}

//获取股票价格
func GetStockPrice() float64 {
	return stockInfo.Data[0].Current
}

//获取股票价格对应的时间
func GetStockLatestTime() string {
	ts := stockInfo.Data[0].Timestamp
	date := component.TimestampToDate_MillSec_Custom(ts, "01-02 15:04")

	return date
}

//组装成股票通知消息
func GetStockInfoNotify() string {
	stockPrice := strconv.FormatFloat(GetStockPrice(), 'f', 2, 64)
	stockLatestTime := GetStockLatestTime()

	notidfy := "众合股价：" + stockPrice + "(截止" + stockLatestTime + ")"

	return notidfy
}

//直接获取一次股票信息
func GetStockPriceOnce() string {
	url := stockBaseUrl + "symbol=" + stockCode
	body, _ := httpHandler.SendHttpGetMessage(url)

	json.Unmarshal(body, &stockInfo)

	return GetStockInfoNotify()
}
