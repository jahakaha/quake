package parser

import (
	"fmt"
	"quake-log-parser/internal/models"
	"regexp"
	"strconv"
	"strings"
)

// Parse data by final result
func ParseData(path string) ([]models.Game, error) {
	data, err := ReadFile(path)
	if err != nil {
		return nil, err
	}
	game, err := SerializeData(data)
	if err != nil {
		return nil, err
	}
	return game, nil
}

// SerializeData func for marshall events by Game struct
func SerializeData(events []models.Event) ([]models.Game, error) {
	fmt.Println("serializeData")
	// starting id with for GameID
	currentGameID := -1
	games := []models.Game{}

	for _, event := range events {
		// fmt.Println(event.EventType)
		// checking for a new game existence
		if event.EventType == models.InitGame {
			// id of game
			currentGameID++
			// create a new game
			game := &models.Game{
				KillMethod: make(map[string]int),
				Players:    map[int]*models.Player{},
			}
			game.ID = currentGameID + 1
			games = append(games, *game)
		} else if event.EventType == models.ClientConnect { /* check fo a new player */
			playerID, err := strconv.Atoi(event.AdditionalData)
			if err != nil {
				return nil, err
			}
			player := games[currentGameID].Players[playerID]
			// if player does not exist we creating new Player
			if player == nil {
				player = &models.Player{}
				player.ID = playerID
			}
			player.ID = playerID
			// add player into Players map[]
			games[currentGameID].Players[playerID] = player
		} else if event.EventType == models.ClientUserinfoChanged { /* check for a player changes info */
			// loop through Players map to get info
			for _, player := range games[currentGameID].Players {
				fmt.Println(player.ID)
				// Split additionalData to get player id
				splittedPlayerID := strings.SplitN(event.AdditionalData, " ", 2)[0]
				playerID, err := strconv.Atoi(splittedPlayerID)
				if err != nil {
					return nil, err
				}
				// fmt.Println("Player id ", playerID)
				// if player exist in map, get name of player from additional data and assign it
				if player.ID == playerID {
					name, err := getPlayerName(event.AdditionalData)
					if err != nil {
						return nil, err
					}
					player.Name = name
				}
			}
		} else if event.EventType == models.Kill { /* check for a new kill */
			// whenever this happen, we increase total kills
			games[currentGameID].TotalKills++
			addKill(event.AdditionalData, &games[currentGameID])
		} else if event.EventType == models.Divider { /* check for a end of game */
			continue
		}
	}
	return games, nil
}

// addKill func to add murder information into game
func addKill(killData string, game *models.Game) (*models.Game, error) {
	// split kill data to get IDs of killer, victim and method of murder
	parts := strings.Split(killData, " ")
	spliitedKillerID := parts[0]
	spliitedVictimID := parts[1]
	// Removing : at the end of kill method
	methodID := strings.Replace(parts[2], ":", "", -1)

	// convert string to int
	killerID, err := strconv.Atoi(spliitedKillerID)
	if err != nil {
		return nil, err
	}
	victimID, err := strconv.Atoi(spliitedVictimID)
	if err != nil {
		return nil, err
	}

	// if  killer is not world and killer does not kill himself
	// then we increasing killer kill count
	if killerID != 1022 && killerID != victimID {
		game.Players[killerID].KillCount++
	}
	// always decresing victim kill count
	game.Players[victimID].KillCount--

	// increase kill method counter
	game.KillMethod[Methods[methodID]]++
	// fmt.Println(methodID, Methods[methodID], game.KillMethod[Methods[methodID]])
	return game, nil
}

// getPlayerName func to regex given string to get only name
func getPlayerName(name string) (string, error) {
	regex := regexp.MustCompile(`n\\([^\\]+)\\t\\`)
	match := regex.FindStringSubmatch(name)
	if len(match) < 2 {
		return "", fmt.Errorf("player name not found in the log entry")
	}
	return match[1], nil
}

// Static map of id and name of kill methods
var Methods map[string]string = map[string]string{
	"0":  models.MOD_UNKNOWN,
	"1":  models.MOD_SHOTGUN,
	"2":  models.MOD_GAUNTLET,
	"3":  models.MOD_MACHINEGUN,
	"4":  models.MOD_GRENADE,
	"5":  models.MOD_GRENADE_SPLASH,
	"6":  models.MOD_ROCKET,
	"7":  models.MOD_ROCKET_SPLASH,
	"8":  models.MOD_PLASMA,
	"9":  models.MOD_PLASMA_SPLASH,
	"10": models.MOD_RAILGUN,
	"11": models.MOD_LIGHTNING,
	"12": models.MOD_BFG,
	"13": models.MOD_BFG_SPLASH,
	"14": models.MOD_WATER,
	"15": models.MOD_SLIME,
	"16": models.MOD_LAVA,
	"17": models.MOD_CRUSH,
	"18": models.MOD_TELEFRAG,
	"19": models.MOD_FALLING,
	"20": models.MOD_SUICIDE,
	"21": models.MOD_TARGET_LASER,
	"22": models.MOD_TRIGGER_HURT,
}
