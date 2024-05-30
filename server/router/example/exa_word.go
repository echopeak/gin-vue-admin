package example

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type WordRouter struct{}

func (e *WordRouter) InitWordRouter(Router *gin.RouterGroup) {
	wordRouter := Router.Group("applet").Use(middleware.OperationRecord())
	wordRouterWithoutRecord := Router.Group("applet")
	exaWordApi := v1.ApiGroupApp.ExampleApiGroup.WordApi
	{
		wordRouter.POST("applet", exaWordApi.CreateWord) // 新建一个单词
	}
	{
		wordRouterWithoutRecord.GET("applet", exaWordApi.GetOneWord) // 获取一个单词
	}
}
