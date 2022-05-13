package models

import (
	"time"

	"github.com/google/uuid"
	pq "github.com/lib/pq"
	"gorm.io/gorm"
)

type GetScheduleHistory struct {
	ID uuid.UUID `gorm:"primaryKey"`
  CreatedAt time.Time
	UpdatedAt time.Time
  DeletedAt gorm.DeletedAt `gorm:"index"`
	ChannelId string `gorm:"index"`
	PersonId string `gorm:"index"`
	Result pq.StringArray `gorm:"type:text[]"`
}

func (getScheduleHistory *GetScheduleHistory) BeforeCreate(tx *gorm.DB) (err error) {
	getScheduleHistory.ID = uuid.New()
	getScheduleHistory.UpdatedAt = time.Now()
	getScheduleHistory.CreatedAt = time.Now()
	
	return
}

func (getScheduleHistory *GetScheduleHistory) BeforeUpdate(tx *gorm.DB) (err error) {
	getScheduleHistory.UpdatedAt = time.Now()

	return
}
