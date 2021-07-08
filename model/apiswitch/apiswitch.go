package apiswitch

import (
	"access-control/internal"
	"access-control/internal/db"
	"time"
)


func TableName() string{
	return internal.InitConfigByViper().TableName
}

type Property struct {
	EffectiveTime time.Time     `json:"effective_time"`
	FailureTime   time.Time     `json:"failure_time"`
	ApiName       string        `json:"api_name"`
	Channel       string        `json:"channel"`
}


func ApiSwitch(p *Property) error{
	return db.GormCoon().Table(TableName()).Create(&p).Error
}

func ApiList(p *Property) (res Property,err error){
	err = db.GormCoon().Table(TableName()).Where("api_name=?",p.ApiName).Find(&res).Error
	return res,err
}