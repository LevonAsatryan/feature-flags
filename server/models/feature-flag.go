package models

import "time"

type FeatureFlag struct {
	ID        string    `json:"id" gorm:"type:uuid;default:gen_random_uuid()"`
	GroupId   string    `json:"groupID" gorm:"not null;uniqueIndex:idx_group_value"`
	Name      string    `json:"name" gorm:"not null;uniqueIndex:idx_group_value"`
	Value     bool      `json:"value" gorm:"not null"`
	CreatedAt time.Time `json:"-" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"-" gorm:"autoUpdateTime"`
}
