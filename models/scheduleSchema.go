package models

import (
	"time"

	"gorm.io/gorm"
)

type Schedule struct {
	gorm.Model
	ChannelId string `gorm:"index"`
	Title string
	Datetime time.Time `gorm:"index"`
	IsTimeSet bool
}
