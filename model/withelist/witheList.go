package withelist

import (
	"access-control/internal/db"
	"unicode"
)

type WitheList struct {
	UserId    int     `json:"user_id"`
	RequestNo string  `json:"request_no"`
	Phone     string  `json:"phone"`
	ID        string  `json:"idNo"`
}

const name = "withe_list"

func WitheCreate(w WitheList) error{
	return db.GormCoon().Table(name).Create(&w).Error
}

func WitheQuery(w WitheList) (res WitheList,err error){
	if w.ID != ""{
		err := db.GormCoon().Table(name).Where("idNo=?",w.ID).Find(&res).Error
		return res,err
	}
	if w.Phone != ""{
		err := db.GormCoon().Table(name).Where("phone=?",w.Phone).Find(&res).Error
		return res,err
	}
	if w.RequestNo != ""{
		err := db.GormCoon().Table(name).Where("request_no=?",w.RequestNo).Find(&res).Error
		return res,err
	}
	if unicode.IsNumber(rune(w.UserId)) != false {
		err := db.GormCoon().Table(name).Where("user_id=?",w.UserId).Find(&res).Error
		return res,err
	}
	return
}