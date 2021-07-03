package builder


type iBuilder interface {
	setWindowType()
	setDoorType()
	setNumFloor()
	getHouse() house
}

func getBuilder(buildType string) iBuilder {
	if buildType == "normal" {
		return createNormalBuilder()
	}

	return nil
}