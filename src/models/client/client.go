package models

import (
	. "config"

	"github.com/jinzhu/gorm"
)

type Connect interface {
	Connect() *gorm.DB
}

type Client struct {
	gorm.Model
	Email string `gorm:"unique_index"`
	Uuid  string
	Key   string
}

func (conn *Client) Connect() *gorm.DB {
	return DatabaseConnection()
}
