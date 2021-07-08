package routes

import (
	"access-control/management"
	"github.com/gin-gonic/gin"
)

func SetRoute(engine *gin.Engine){
	r := engine.Group("access-control")

	apiRoute := r.Group("/v1")
	{
		apiRoute.POST("/createApi",management.ApiCreate)
		apiRoute.POST("/RateLimit",management.ApiRateLimit)
	}
}