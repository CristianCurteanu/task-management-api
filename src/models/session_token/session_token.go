package models

import (
	. "config"
	"time"

	"github.com/jinzhu/gorm"
)

type Connect interface {
	Connect() *gorm.DB
}

type SessionToken struct {
	gorm.Model
	Token      string `gorm:"unique_index"`
	ExpirestAt time.Time
}

func (conn *SessionToken) Connect() *gorm.DB {
	return DatabaseConnection()
}
