package main

import (
	"net/http"

	_ "lkserver/model/accessTokenManager"
	_ "lkserver/model/crontab"

	"lkserver/model/wecomMsgManager"

	"github.com/gin-gonic/gin"
)

func main() {
	//创建一个默认的路由引擎
	app := gin.Default()

	//注册时间戳转日期的函数,必须放在模板前面
	// app.SetFuncMap(template.FuncMap{
	// 	"TimestampToDate": TimestampToDate,
	// })

	//加载html模板
	//app.LoadHTMLGlob("templates/**/*")

	//配置静态web目录 第一个参数表示路由，第二个参数表示映射的目录
	//app.Static("/static", "../static")

	//设置 /
	app.GET("/lifeapp/lkbd/send/:user/:cont", func(c *gin.Context) {
		user := c.Param("user")
		cont := c.Param("cont")
		var myMsg wecomMsgManager.PostMsg
		myMsg.ToUser = user
		myMsg.AgentId = "1000003"
		myMsg.MsgType = "text"
		myMsg.Text = wecomMsgManager.Message{Cont: cont}
		myMsg.DuplicateCheckInterval = "600"

		status, err := wecomMsgManager.WecomPostMessage(1, myMsg)
		if err != nil {
			status = err.Error()
		}

		c.JSON(http.StatusOK, gin.H{
			"receiver": user,
			"status":   status,
			"content":  cont,
		})
	})

	app.GET("/lifeapp", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "good",
		})
	})

	// stockPrice.GetStockPriceOnce()
	// oilPrice.GetStockPriceOnce()
	// collector.CollectInfoToPush()

	// 启动HTTP服务，默认在0.0.0.0:8080
	app.Run(":80")
}
