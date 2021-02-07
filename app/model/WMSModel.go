package model

import (
	"time"
)

type WMSModel struct {
	Id      int `json:"id" gorm"primarykey"`
	Created time.Time `json:"created" gorm:"default:null"`
	Updated time.Time `json:"updated" gorm:"default:null"`
}