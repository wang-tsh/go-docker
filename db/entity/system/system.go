package system

import "go-docker/db/entity"

var tableName = "DE_ACCESS_SYSTEM"

func (AccessSystem) TableName() string {
	return tableName
}

type AccessSystem struct {
	entity.Base
	ID           int32  `json:"id"           gorm:"primaryKey;column:ID"`
	SystemName   string `json:"systemName"   gorm:"column:SYSTEM_NAME"`
	SystemCode   string `json:"systemCode"   gorm:"column:SYSTEM_CODE"`
	Introduction string `json:"introduction" gorm:"column:INTRODUCTION"`
	NetType      int    `json:"netType"      gorm:"column:NET_TYPE"`
	OrgCode      string `json:"orgCode"      gorm:"column:ORG_CODE"`
}
