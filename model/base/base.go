package base

type BaseModel struct {
	Id              int            `json:"id" gorm:"column:id"`
	CreateTime      string         `json:"create_time" gorm:"column:create_time"`
	UpdateTime      string         `json:"update_time" gorm:"column:update_time"`
}
