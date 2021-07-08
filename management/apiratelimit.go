package management

import (
	"access-control/internal"
	"access-control/internal/ratelimit"
	"access-control/model/apiswitch"
	"github.com/gin-gonic/gin"
	"github.com/wonderivan/logger"
	"net/http"
	"time"
)

func ApiRateLimit(ctx *gin.Context){
	var (
		tb  *ratelimit.TokenBucket
		err error
		res  apiswitch.Property
	)

	if err = ctx.ShouldBind(&res);err != nil{
		ctx.JSON(http.StatusBadRequest,gin.H{"code":"0","msg":"格式错误"})
		return
	}

	res,err = apiswitch.ApiList(&res)

	if err != nil{
		logger.Info("api未配置")
	}else {
		if  res.EffectiveTime.Unix() < time.Now().Unix(){
			ctx.JSON(http.StatusBadRequest,gin.H{"code":"0","msg":"该时间段禁止访问"})
			return
		}
	}

	tb = ratelimit.NewTokenBucket(internal.InitConsByViper().Capacity, internal.InitConsByViper().ProdTokenIntervalMs, internal.InitConsByViper().ProdTokenNumEveryInterval)
	err = tb.TryAquire()
	if err != nil{
		tb.WaitUntilAquire()
	}

	ctx.JSON(http.StatusOK,gin.H{"code":0,"msg":"请求成功"})

}

