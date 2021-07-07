package flyweight

type game struct {
	terrorists []*player
	counterTerrorists []*player
}

func newGame() *game {
	return &game{
		terrorists: make([]*player, 1),
		counterTerrorists: make([]*player, 1),
	}
}

func (g *game) addTerrorists(dressType string) {
	player := newPlayer("T", dressType)
	g.terrorists = append(g.terrorists, player)
	return
}

func (g *game) addCounterTerrorists(dressType string) {
	player := newPlayer("CT", dressType)
	g.terrorists = append(g.counterTerrorists, player)
	return
}