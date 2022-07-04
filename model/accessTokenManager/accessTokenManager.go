package accessTokenManager

/*
 * 该模块用于保持最新的wecomToken，并降低向服务器获取wecomToken的次数
 * 每个Token有效期为2小时
 */

import (
	"fmt"

	"github.com/robfig/cron/v3"

	"lkserver/model/httpHandler"
	"lkserver/model/wecomCfgReader"
)

type AppCronInfo struct {
	AppName     string
	AccessToken string
	CronSpec    string
	CronObj     *cron.Cron
	CronEntryId cron.EntryID
}

var (
	AppCronInfoArr []AppCronInfo
)

func init() {
	AppCronInfoArr = make([]AppCronInfo, wecomCfgReader.MaxIdx)

	for i := 0; i < wecomCfgReader.MaxIdx; i++ {
		go updateTokenFirst(i)
	}

	for i := 0; i < wecomCfgReader.MaxIdx; i++ {
		AppCronInfoArr[i].AppName = wecomCfgReader.GetAppInfo(i).Name
		AppCronInfoArr[i].AccessToken = ""
		AppCronInfoArr[i].CronSpec = "@every 2h"
		AppCronInfoArr[i].CronObj = &cron.Cron{}
		AppCronInfoArr[i].CronEntryId = cron.EntryID(0)
	}
}

//周期性更新Token
func UpdateTokenCycle(appIdx int) {
	if latestToken, err := getAccessTokenOnce(appIdx); err == nil {
		SetLatestWecomToken(appIdx, latestToken)
	} else {
		fmt.Println(err)
	}
}

//用于触发更新Token的goroutine，更新Token成功后更新cron对象，使其重新开始倒计时
func updateTokenFirst(appIdx int) {
	//获取并更新Token
	if latestToken, err := getAccessTokenOnce(appIdx); err == nil {
		SetLatestWecomToken(appIdx, latestToken)
		fmt.Println(latestToken)
	} else {
		fmt.Println(err)
	}

	c := cron.New()
	AppCronInfoArr[appIdx].CronObj = c
	AppCronInfoArr[appIdx].CronEntryId, _ = c.AddFunc(AppCronInfoArr[appIdx].CronSpec, func() { UpdateTokenCycle(appIdx) })
	c.Start()
}

func GetAccessToken(AppIdx int) string {
	token := ""

	if AppIdx < wecomCfgReader.MaxIdx {
		token = AppCronInfoArr[AppIdx].AccessToken
	}

	return token
}

func SetLatestWecomToken(appIdx int, latestToken string) {
	if AppCronInfoArr[appIdx].AccessToken != latestToken {
		AppCronInfoArr[appIdx].AccessToken = latestToken
	}
}

func getAccessTokenOnce(appIdx int) (string, error) {
	token, err := httpHandler.GetToken(appIdx)
	return token, err
}

//外部失败后触发更新Token
func UpdateTokenWhenFail(appIdx int) {
	c := AppCronInfoArr[appIdx].CronObj
	c.Stop()
	c.Remove(AppCronInfoArr[appIdx].CronEntryId)
	AppCronInfoArr[appIdx].CronEntryId, _ = c.AddFunc(AppCronInfoArr[appIdx].CronSpec, func() { UpdateTokenCycle(appIdx) })
	c.Start()
}
