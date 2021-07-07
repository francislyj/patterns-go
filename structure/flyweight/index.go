package flyweight

import "fmt"

func TestFlyWeight() {
	game := newGame()

	game.addTerrorists(TerroristDressType)
	game.addTerrorists(TerroristDressType)
	game.addTerrorists(TerroristDressType)
	game.addTerrorists(TerroristDressType)

	game.addCounterTerrorists(CounterTerroristDressType)
	game.addCounterTerrorists(CounterTerroristDressType)
	game.addCounterTerrorists(CounterTerroristDressType)

	for dressType, dress := range getDressFactory().dressMap {
		fmt.Printf("DressColorType: %s\nDressColor: %s\n", dressType, dress.getColor())
	}
}
