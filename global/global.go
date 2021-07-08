package global

import (
	"access-control/config"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)


var (
	GVA_DB     *gorm.DB
	GVA_CONFIG config.Server
	GVA_VP     *viper.Viper
)
