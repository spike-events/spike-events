package models

import (
	"encoding/json"
	"gorm.io/datatypes"
	"gorm.io/gorm"
	"spike.io/bin"
)

type Message struct {
	Base
	MessageData datatypes.JSON
	Message     *bin.Message `gorm:"-"`
}

func (u *Message) BeforeSave(tx *gorm.DB) (err error) {
	payload, err := json.Marshal(u.Message)
	if err != nil {
		return
	}
	u.MessageData = payload
	return nil
}

// AfterFind parse json deal
func (u *Message) AfterFind(tx *gorm.DB) (err error) {
	err = json.Unmarshal(u.MessageData, &u.Message)
	u.Message.Offset = u.ID
	return
}
