package management

import (
	"access-control/model/withelist"
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/wonderivan/logger"
)

func WitheLimit(ctx *gin.Context){
	var (
		w withelist.WitheList
		err error
	)
	if err = ctx.ShouldBind(&w);err != nil{
		ctx.JSON(http.StatusBadRequest,gin.H{"code":"0","msg":"参数错误"})
		return
	}

	w,err = withelist.WitheQuery(w)
	if err != nil{
		logger.Info("白名单未配置")
		ctx.JSON(http.StatusOK,gin.H{"code":"0","msg":"白名单未配置，禁止准入"})
		return
	}else {
		ctx.JSON(http.StatusOK,gin.H{"code":"0","msg":"请求成功"})
		return
	}

}
