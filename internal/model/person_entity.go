package model

import (
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

type PersonEntity struct {
	Id       uuid.UUID `gorm:"primarykey,type:uuid;default:uuid_generate_v4()"`
	Username string    
	Name     string    
	Birthday time.Time `gorm:"type:date"`
	Stack    pq.StringArray  `gorm:"type:text[]"`
}

func (PersonEntity) TableName() string {
    return "people"
}