package httpHandler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"lkserver/model/wecomCfgReader"
	"net/http"
)

type AccessToken struct {
	Errcode     int    `json:"errcode"`
	Errmsg      string `json:"errmsg"`
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
}

type SendMsgSucc struct {
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
	MsgID   string `json:"msgid"`
}

type SendMsgFail struct {
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}

func packAccessTokenUrl(appIdx int) string {
	return wecomCfgReader.GetAccessTokenUrl() + "corpid=" + wecomCfgReader.GetCorpId() + "&corpsecret=" + wecomCfgReader.GetAppInfo(appIdx).Token
}

func GetToken(appIdx int) (string, error) {
	getTokenUrl := packAccessTokenUrl(appIdx)
	resp, err := http.Get(getTokenUrl)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	token := AccessToken{}
	json.Unmarshal(body, &token)

	return token.AccessToken, err
}

func SendHttpPostMessage(url string, payload []byte) (string, error) {
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(payload))
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	//读取响应体
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(body))

	return string(body), err
}

//标准的http请求，返回响应体
func SendHttpGetMessage(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	//读取响应体
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	return body, err
}
