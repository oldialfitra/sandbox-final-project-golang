package models

import (
	"sandbox-final-project/server/api/security"
	"time"

	_ "github.com/jinzhu/gorm"
)

type User struct {
	ID       int       `gorm:"primary_key;auto_increment" json:"id"`
	Username string    `gorm:"size:255;not null;unique" json:"username"`
	Password string    `gorm:"size:100;not null;" json:"password"`
	Created  time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
}

func (u *User) BeforeSave() error {
	hashedPassword, err := security.Hash(u.Password)
	if err != nil {
		return err
	}
}
