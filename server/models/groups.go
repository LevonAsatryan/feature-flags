package models

import "time"

type Group struct {
	ID        string    `json:"id" gorm:"type:uuid;default:gen_random_uuid()"`
	Name      string    `json:"name" gorm:"unique";not null`
	CreatedAt time.Time `json:"-" gorm:"<-:create"`
	UpdatedAt time.Time `json:"-"`
}
