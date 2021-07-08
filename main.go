package main

import (
	"access-control/routes"
	"github.com/gin-gonic/gin"
	"github.com/wonderivan/logger"
)



func main() {
	logger.SetLogger("./log.json")
	r := gin.Default()
	routes.SetRoute(r)
	r.Run()
}
