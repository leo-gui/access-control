package management

import "C"
import (
	"access-control/model/channel"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ChannelCreate(ctx *gin.Context){
	var (
		err error
		c channel.Channel
	)
	if err = ctx.ShouldBind(&c);err != nil{
		ctx.JSON(http.StatusBadRequest,gin.H{"code":"0","msg":"参数错误"})
		return
	}
	if err = channel.ChannelCreate(&c);err != nil{
		ctx.JSON(http.StatusInternalServerError,gin.H{"code":10086,"msg":"创建失败"})
		return
	}
	ctx.JSON(http.StatusOK,gin.H{"code":0,"msg":"创建成功"})
}