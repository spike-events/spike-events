package models

import (
	"encoding/json"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type GroupID string

type Topic struct {
	Base
	Table      string
	Name       string `gorm:"index"`
	OffsetData datatypes.JSON
	Offset     map[GroupID]*int64 `gorm:"-"`
}

func (u *Topic) BeforeSave(tx *gorm.DB) (err error) {
	payload, err := json.Marshal(u.OffsetData)
	if err != nil {
		return
	}
	u.OffsetData = payload
	return nil
}

// AfterFind parse json deal
func (u *Topic) AfterFind(tx *gorm.DB) (err error) {
	err = json.Unmarshal(u.OffsetData, &u.Offset)
	return
}
