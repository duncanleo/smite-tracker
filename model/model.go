package model

import "time"

// Player represents a player
type Player struct {
	ID        uint       `gorm:"primary_key"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `sql:"index" json:"-"`
	Name      string     `gorm:"name" json:"name"`
}

// PlayerData represents the state of a player at a given time
type PlayerData struct {
	ID        uint       `gorm:"primary_key"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `sql:"index" json:"-"`
	Player    Player     `json:"player"`
	PlayerID  uint       `json:"-"`
}
