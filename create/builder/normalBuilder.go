package builder

type normalBuilder struct {
	windowType string
	doorType string
	floor int
}

func createNormalBuilder() iBuilder {
	return &normalBuilder{}
}

func (n *normalBuilder) setWindowType() {
	n.windowType = "normal window"
}

func (n *normalBuilder) setDoorType() {
	n.doorType = "normal door"
}

func (n *normalBuilder) setNumFloor() {
	n.floor = 20
}

func (n *normalBuilder) getHouse() house {
	return house{
		windowType: n.windowType,
		doorType: n.doorType,
		floor: n.floor,
	}
}