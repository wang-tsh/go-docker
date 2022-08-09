package entity

import (
	"fmt"
	"time"
)

type Base struct {
	Creator    string   `json:"creator"           gorm:"column:CREATOR"`
	Modifier   string   `json:"modifier"          gorm:"column:MODIFIER"`
	CreateTime DateTime `json:"createTime"        gorm:"column:CREATE_TIME;autoCreateTime"`
	UpdateTime DateTime `json:"updateTime"        gorm:"column:UPDATE_TIME;autoUpdateTime"`
}

type DateTime time.Time

func (t DateTime) MarshalJSON() ([]byte, error) {
	var stamp = fmt.Sprintf("\"%s\"", time.Time(t).Format("2006-01-02 15:04:05"))
	return []byte(stamp), nil
}

//func (t Base) MarshalJSON() ([]byte, error) {
//	type TmpJSON Base
//	return json.Marshal(&struct {
//		TmpJSON
//		CreateTime DateTime `json:"createTime"`
//		UpdateTime DateTime `json:"updateTime"`
//	}{
//		TmpJSON:    (TmpJSON)(t),
//		CreateTime: DateTime(t.CreateTime),
//		UpdateTime: DateTime(t.UpdateTime),
//	})
//}
