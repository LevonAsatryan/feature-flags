package models

import "time"

type Group struct {
	ID        string    `gorm:"type:uuid;default:gen_random_uuid()"`
	Name      string    `json:"name" gorm:"unique";not null`
	CreatedAt time.Time `gorm:"<-:create"`
	UpdatedAt time.Time
}
