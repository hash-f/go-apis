package game

type GameStore map[string](*Game)

func NewGameStore() GameStore {
	return make(map[string](*Game), 0)
}

func (gs *GameStore) AddFakeGame() {
	g := NewGame()
	g.Id = "1"
	g.word = "abcd"
	(*gs)["1"] = g
}
