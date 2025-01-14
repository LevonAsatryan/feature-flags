package models

import "time"

type FeatureFlag struct {
	ID        string    `gorm:"type:uuid;default:gen_random_uuid()"`
	GroupId   string    `josn:"groupID" gorm:"not null;uniqueIndex:idx_group_value"`
	Name      string    `json:"name" gorm:"not null;uniqueIndex:idx_group_value"`
	Value     bool      `json:"value" gorm:"not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}
