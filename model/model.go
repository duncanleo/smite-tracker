package model

import "time"

// Player represents a player
type Player struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `sql:"index" json:"-"`
	Name      string     `gorm:"name" json:"name"`
	GameID    int        `gorm:"game_id" json:"game_id"`
	Status    string     `json:"status"`
}

// PlayerData represents the state of a player at a given time
type PlayerData struct {
	ID                uint       `gorm:"primary_key" json:"id"`
	CreatedAt         time.Time  `json:"created_at"`
	UpdatedAt         time.Time  `json:"updated_at"`
	DeletedAt         *time.Time `sql:"index" json:"-"`
	Player            Player     `json:"-"`
	PlayerID          uint       `json:"-"`
	HoursPlayed       int        `json:"hours_played"`
	Leaves            int        `json:"leaves"`
	Level             int        `json:"level"`
	Losses            int        `json:"losses"`
	MasteryLevel      int        `json:"mastery_level"`
	Wins              int        `json:"wins"`
	TotalAchievements int        `json:"Total_Achievements"`
	TotalWorshippers  int        `json:"Total_Worshippers"`
	TopGod            God        `json:"top_god"`
	TopGodID          uint       `json:"-"`
	TopGodRank        int        `json:"top_god_rank"`
	TopGodWorshippers int        `json:"top_god_worshippers"`
}

// God represents a god in SMITE
type God struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `sql:"index" json:"-"`
	GameID    string     `gorm:"game_id" json:"game_id"`
	Name      string     `gorm:"name" json:"name"`
}
