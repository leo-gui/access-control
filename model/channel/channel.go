package channel

import (
	"access-control/internal"
	"access-control/internal/db"
	"access-control/model/apiswitch"
	"time"
)

type Channel struct {
	ChannelName   string `json:"channel_name"`
	IsSwitch      bool   `json:"is_switch"`
	EffectiveTime time.Time     `json:"effective_time"`
	FailureTime   time.Time     `json:"failure_time"`
}



func TableName() string{
	return internal.InitConfigByViper().TableName
}

func ChannelCreate(c *Channel) error{
	return db.GormCoon().Table(TableName()).Create(&c).Error
}

func ChannelList(p *apiswitch.Property) (res Channel,err error){
	err = db.GormCoon().Table(TableName()).Where("channel_name=?",p.Channel).Find(&res).Error
	return res,err
}
