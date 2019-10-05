package brawl

// Player represents a player in Brawl Stars
type Player struct {
	ThreeVs3Victories    int       `json:"3vs3Victories"`
	BestRoboRumbleTime   int       `json:"bestRoboRumbleTime"`
	BestTimeAsBigBrawler int       `json:"bestTimeAsBigBrawler"`
	Brawlers             []Brawler `json:"brawlers"`
	Club                 struct {
		Name string `json:"name"`
		Tag  string `json:"tag"`
	} `json:"club"`
	DuoVictories    int    `json:"duoVictories"`
	ExpLevel        int    `json:"expLevel"`
	ExpPoints       int    `json:"expPoints"`
	HighestTrophies int    `json:"highestTrophies"`
	Name            string `json:"name"`
	NameColor       string `json:"nameColor"`
	SoloVictories   int    `json:"soloVictories"`
	Tag             string `json:"tag"`
	Trophies        int    `json:"trophies"`
}

// Brawler represents a brawler in Brawl Stars
type Brawler struct {
	HighestTrophies int    `json:"highestTrophies"`
	ID              int    `json:"id"`
	Name            string `json:"name"`
	Power           int    `json:"power"`
	Rank            int    `json:"rank"`
	StarPowers      []struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"starPowers"`
	Trophies int `json:"trophies"`
}
