package models

type Event struct {
	Timestamp      string
	EventType      string
	AdditionalData string
}

type Game struct {
	ID         int
	TotalKills int
	Players    map[int]*Player
	KillMethod map[string]int
}

type Player struct {
	ID         int
	Name       string
	KillCount  int
	DeathCount int
}

const (
	InitGame              = "InitGame"
	ClientConnect         = "ClientConnect"
	ClientUserinfoChanged = "ClientUserinfoChanged"
	Kill                  = "Kill"
	Divider               = "------------------------------------------------------------"
)
