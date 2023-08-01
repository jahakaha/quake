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

const (
	MOD_UNKNOWN        = "MOD_UNKNOWN"
	MOD_SHOTGUN        = "MOD_SHOTGUN"
	MOD_GAUNTLET       = "MOD_GAUNTLET"
	MOD_MACHINEGUN     = "MOD_MACHINEGUN"
	MOD_GRENADE        = "MOD_GRENADE"
	MOD_GRENADE_SPLASH = "MOD_GRENADE_SPLASH"
	MOD_ROCKET         = "MOD_ROCKET"
	MOD_ROCKET_SPLASH  = "MOD_ROCKET_SPLASH"
	MOD_PLASMA         = "MOD_PLASMA"
	MOD_PLASMA_SPLASH  = "MOD_PLASMA_SPLASH"
	MOD_RAILGUN        = "MOD_RAILGUN"
	MOD_LIGHTNING      = "MOD_LIGHTNING"
	MOD_BFG            = "MOD_BFG"
	MOD_BFG_SPLASH     = "MOD_BFG_SPLASH"
	MOD_WATER          = "MOD_WATER"
	MOD_SLIME          = "MOD_SLIME"
	MOD_LAVA           = "MOD_LAVA"
	MOD_CRUSH          = "MOD_CRUSH"
	MOD_TELEFRAG       = "MOD_TELEFRAG"
	MOD_FALLING        = "MOD_FALLING"
	MOD_SUICIDE        = "MOD_SUICIDE"
	MOD_TARGET_LASER   = "MOD_TARGET_LASER"
	MOD_TRIGGER_HURT   = "MOD_TRIGGER_HURT"
)
