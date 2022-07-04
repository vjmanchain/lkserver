package crontab

import (
	"lkserver/model/lifeNotify/collector"
	"lkserver/model/lifeNotify/oilPrice"
	"lkserver/model/lifeNotify/stockPrice"

	"github.com/robfig/cron/v3"
)

func init() {
	c := cron.New()

	//周一至周五早上8:20 更新油价
	c.AddFunc("20 8 * * 1-5", oilPrice.GetOilInfoFromTianxing)
	//周一至周五早上8:20 更新股价
	c.AddFunc("20 8 * * 1-5", stockPrice.GetStockInfoFromXueqiu)

	//周一至周五早上8:30 向所有用户推送
	c.AddFunc("30 8 * * 1-5", collector.CollectInfoToPush)

	c.Start()
}
