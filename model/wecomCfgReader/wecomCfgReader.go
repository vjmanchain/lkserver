package wecomCfgReader

/*
 * 该模块用于读取配置，外部模块需调用接口获取配置
 */

import (
	"strconv"

	"github.com/spf13/viper"
)

const (
	WorkIdx = iota
	LifeIdx
	MaxIdx
)

type App struct {
	Name  string
	Id    string
	Token string
}

var (
	apps           []App
	corpId         string
	accessTokenUrl string
	wecomMsgUrl    string
)

func init() {
	v := viper.New()
	v.AddConfigPath("config")
	v.SetConfigName("wecom")
	v.SetConfigType("yml")

	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}

	for i := 0; i < MaxIdx; i++ {
		app := App{
			Name:  v.GetString("Apps." + strconv.Itoa(i) + ".name"),
			Id:    v.GetString("Apps." + strconv.Itoa(i) + ".id"),
			Token: v.GetString("Apps." + strconv.Itoa(i) + ".token"),
		}
		apps = append(apps, app)
	}

	corpId = v.GetString("CorpId")
	accessTokenUrl = v.GetString("AccessTokenUrl")
	wecomMsgUrl = v.GetString("WecomMsgUrl")
}

func GetWorkAppInfo() App {
	return apps[WorkIdx]
}

func GetLifeAppInfo() App {
	return apps[LifeIdx]
}

func GetAppInfo(appIdx int) App {
	return apps[appIdx]
}

func GetCorpId() string {
	return corpId
}

func GetAccessTokenUrl() string {
	return accessTokenUrl
}

func GetWecomMsgUrl() string {
	return wecomMsgUrl
}
