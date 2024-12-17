package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Group struct {
	gorm.Model
	Id   uuid.UUID `gorm:"type:uuid;default:gen_random_uuid()"`
	Name string
}
