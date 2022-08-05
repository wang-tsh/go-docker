package system

import "go-docker/db/entity"

var TableName = "DE_ACCESS_SYSTEM"

type AccessSystem struct {
	entity.Base
	ID           int32  `json:"id"           gorm:"primaryKey;column:ID"`
	SystemName   string `json:"systemName"   gorm:"systemName"`
	SystemCode   string `json:"systemCode"   gorm:"systemCode"`
	Introduction string `json:"introduction" gorm:"introduction"`
	NetType      int32  `json:"netType"      gorm:"netType"`
	OrgCode      string `json:"orgCode"      gorm:"orgCode"`
}
