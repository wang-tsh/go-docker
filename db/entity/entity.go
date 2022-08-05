package entity

import "time"

type Base struct {
	Creator    string    `json:"creator"           gorm:"creator"`
	Modifier   string    `json:"modifier"          gorm:"modifier"`
	CreateTime time.Time `json:"createTime"        gorm:"createTime"`
	UpdateTime time.Time `json:"updateTime"        gorm:"updateTime"`
}
