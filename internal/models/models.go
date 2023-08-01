package models

// Event represents a game event with its timestamp, event type, and additional data.
type Event struct {
	Timestamp      string // Timestamp of the event
	EventType      string // Type of the event (e.g., InitGame, Kill, ClientConnect, etc.)
	AdditionalData string // Additional data associated with the event
}

// Game represents a game with its ID, total number of kills, players, and kill methods.
type Game struct {
	ID         int             // Unique ID of the game
	TotalKills int             // Total number of kills in the game
	Players    map[int]*Player // Map of player IDs to their respective Player structs
	KillMethod map[string]int  // Map of kill methods and their corresponding counts
}

// Player represents a player with their ID, name, kill count, and death count.
type Player struct {
	ID         int    // Unique ID of the player
	Name       string // Name of the player
	KillCount  int    // Number of kills for the player
	DeathCount int    // Number of deaths for the player
}

// Response represents the response containing game data.
type Response struct {
	Game map[string]*GameReponse // Map of game IDs to their respective GameReponse structs
}

// GameReponse represents the response data for a game, including total kills, player names, and kill counts.
type GameReponse struct {
	TotalKills int            // Total number of kills in the game
	Players    []string       // List of player names in the game
	Kills      map[string]int // Map of player names to their respective kill counts
	KillMeans  map[string]int // Map of kill methods to their corresponding counts
}

// Constants representing various event types.
const (
	InitGame              = "InitGame"
	ClientConnect         = "ClientConnect"
	ClientUserinfoChanged = "ClientUserinfoChanged"
	Kill                  = "Kill"
	Divider               = "------------------------------------------------------------"
)

// Constants representing various kill methods.
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
