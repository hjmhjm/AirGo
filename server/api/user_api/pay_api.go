package user_api

import (
	"github.com/gin-gonic/gin"
	"github.com/ppoonk/AirGo/global"
	"github.com/ppoonk/AirGo/model"
	"github.com/ppoonk/AirGo/service"
	"github.com/ppoonk/AirGo/utils/response"
)

// 获取已激活支付列表
func GetEnabledPayList(ctx *gin.Context) {
	list, _, err := service.CommonSqlFind[model.Pay, string, []model.Pay]("status = true")
	if err != nil {
		global.Logrus.Error(err.Error())
		response.Fail(err.Error(), nil, ctx)
		return
	}
	response.OK("GetEnabledPayList success", list, ctx)
}
