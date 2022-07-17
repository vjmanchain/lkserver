package collector

import (
	"fmt"
	"lkserver/model/lifeNotify/oilPrice"
	"lkserver/model/lifeNotify/stockPrice"
	"lkserver/model/wecomCfgReader"
	"lkserver/model/wecomMsgManager"
)

//搜集推送的信息
func CollectInfoToPush() {
	oilInfo := oilPrice.GetOilInfoNotify()
	stockInfo := stockPrice.GetStockInfoNotify()

	message := "尊贵的列控老伙计，早上好！" + "\n\n" + oilInfo + "\n" + stockInfo + "\n\n" + "记得72小时核酸检测哦~\n\n更多功能正在开发中，敬请期待！"

	var myMsg wecomMsgManager.PostMsg
	myMsg.ToUser = "@all"
	myMsg.AgentId = "1000003"
	myMsg.MsgType = "text"
	myMsg.Text = wecomMsgManager.Message{Cont: message}
	myMsg.DuplicateCheckInterval = "600"

	status, err := wecomMsgManager.WecomPostMessage(wecomCfgReader.LifeIdx, myMsg)
	if err != nil {
		status = err.Error()
	}
	fmt.Println(status)
}
