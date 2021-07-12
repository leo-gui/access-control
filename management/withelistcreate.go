package management

import (
	"access-control/model/withelist"
	"github.com/gin-gonic/gin"
	"net/http"
)

func WitheCreate(ctx *gin.Context){
	var (
		err error
		w withelist.WitheList
	)
	if err = ctx.ShouldBind(&w);err != nil{
		ctx.JSON(http.StatusBadRequest,gin.H{"code":"0","msg":"参数错误"})
		return
	}
	if err = withelist.WitheCreate(w);err != nil{
		ctx.JSON(http.StatusInternalServerError,gin.H{"code":10086,"msg":"创建失败"})
		return
	}
	ctx.JSON(http.StatusOK,gin.H{"code":0,"msg":"创建成功"})
}
