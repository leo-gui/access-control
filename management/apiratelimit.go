package management

import (
	"access-control/internal"
	"access-control/internal/ratelimit"
	"access-control/model/apiswitch"
	"access-control/model/channel"
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
		data channel.Channel
	)

	if err = ctx.ShouldBind(&res);err != nil{
		ctx.JSON(http.StatusBadRequest,gin.H{"code":"0","msg":"参数错误"})
		return
	}

	//根据时间段来过滤
	res,err = apiswitch.ApiList(&res)
	if err != nil{
		logger.Info("api未配置")
	}else {
		if  res.EffectiveTime.Unix() < time.Now().Unix(){
			ctx.JSON(http.StatusBadRequest,gin.H{"code":"0","msg":"该时间段禁止访问"})
			return
		}
	}

	//根据渠道来过滤
	data,err = channel.ChannelList(&res)
	if err != nil{
		logger.Info("channel未配置")
	}else {
		if data.IsSwitch == false{
			ctx.JSON(http.StatusBadRequest,gin.H{"code":"0","msg":"该渠道禁止访问"})
			return
		}
	}

	//限流
	tb = ratelimit.NewTokenBucket(internal.InitConsByViper().Capacity, internal.InitConsByViper().ProdTokenIntervalMs, internal.InitConsByViper().ProdTokenNumEveryInterval)
	err = tb.TryAquire()
	if err != nil{
		tb.WaitUntilAquire()
	}

	ctx.JSON(http.StatusOK,gin.H{"code":0,"msg":"请求成功"})

}

