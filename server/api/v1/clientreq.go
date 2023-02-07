package v1

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"jjchatgpt/global"
	"jjchatgpt/model"
	"jjchatgpt/model/response"
	"jjchatgpt/service"
)

func Questansw(c *gin.Context) {
	var a model.ClientQuest
	_ = c.ShouldBind(&a)

	if respstr, err := service.RespClientQuest(a.Prompt); err != nil {
		global.Chatgptcn_LOG.Error("失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败:"+err.Error(), c)
	} else {
		response.OkWithData(gin.H{"answ": respstr}, c)
	}
}
