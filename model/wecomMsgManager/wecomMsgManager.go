package wecomMsgManager

import (
	"encoding/json"
	"fmt"
	"lkserver/model/accessTokenManager"
	"lkserver/model/httpHandler"
	"lkserver/model/wecomCfgReader"
)

type Message struct {
	Cont string `json:"content"`
}

type PostMsg struct {
	ToUser                 string  `json:"touser"`
	AgentId                string  `json:"agentid"`
	MsgType                string  `json:"msgtype"`
	Text                   Message `json:"text"`
	DuplicateCheckInterval string  `json:"duplicate_check_interval"`
}

func packSendMessageUrl(appIdx int) string {
	return wecomCfgReader.GetWecomMsgUrl() + "access_token=" + accessTokenManager.GetAccessToken(appIdx)
}

func WecomPostMessage(appIdx int, msg PostMsg) (string, error) {
	url := packSendMessageUrl(appIdx)
	postMsg, _ := json.Marshal(msg)
	_, err := httpHandler.SendHttpPostMessage(url, postMsg)
	if err != nil {
		fmt.Println(err)
		accessTokenManager.UpdateTokenWhenFail(appIdx)
	}
	return "Send Success", err
}
