package management

import (
	"access-control/model/apiswitch"
	"github.com/gin-gonic/gin"
	"net/http"
)


func ApiCreate(ctx *gin.Context){
	var (
		err error
		property apiswitch.Property
	)
	if err = ctx.ShouldBind(&property);err != nil{
		ctx.JSON(http.StatusBadRequest,gin.H{"code":"-1","msg":err})
		return
	}
	if err = apiswitch.ApiSwitch(&property);err != nil{
		ctx.JSON(http.StatusInternalServerError,gin.H{"code":10086,"msg":"创建失败"})
		return
	}
	ctx.JSON(http.StatusOK,gin.H{"code":0,"msg":"创建成功"})
}
