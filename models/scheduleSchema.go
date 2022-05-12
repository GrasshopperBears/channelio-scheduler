package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Schedule struct {
	ID uuid.UUID `gorm:"primaryKey"`
  CreatedAt time.Time
  DeletedAt gorm.DeletedAt `gorm:"index"`
	ChannelId string `gorm:"index"`
	Title string
	Datetime time.Time `gorm:"index"`
	IsTimeSet bool
}

func (schedule *Schedule) BeforeCreate(tx *gorm.DB) (err error) {
	schedule.ID = uuid.New()
	schedule.CreatedAt = time.Now()

	return
}
