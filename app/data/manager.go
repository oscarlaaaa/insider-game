package data

// Manager manages all of the game sessions occurring
type Manager struct {
	GameIDs map[int]*GameRoom
	Games   map[string]*GameRoom
}

var GM Manager = Manager{
	GameIDs: make(map[int]*GameRoom),
	Games:   make(map[string]*GameRoom),
}

// Create new game session
func (gm Manager) addGame(game *GameRoom) {

	gm.GameIDs[game.id] = game
	gm.Games[game.name] = game

}
