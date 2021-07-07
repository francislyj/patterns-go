package flyweight

import "fmt"

const (
	TerroristDressType = "tDress"
	CounterTerroristDressType = "ctDress"
)

var (
	dressFactoryInstance = &dressFactory{
		dressMap: make(map[string]dress),
	}
)

type dressFactory struct {
	dressMap map[string]dress
}

func (d *dressFactory) getDressByType(dressType string) (dress, error) {

	value, exists := d.dressMap[dressType]
	if exists {
		return value, nil
	}

	if dressType == TerroristDressType {
		d.dressMap[dressType] = newTerroristDress("black")
		return d.dressMap[dressType], nil
	}

	if dressType == CounterTerroristDressType {
		d.dressMap[dressType] = newCounterTerroristDress("blue")
		return d.dressMap[dressType], nil
	}

	return nil, fmt.Errorf("type is wrong")

}

func getDressFactory() *dressFactory {
	return dressFactoryInstance
}