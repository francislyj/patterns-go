package builder


type director struct {
	builder iBuilder
}

func newDirector(builder iBuilder) *director {
	return &director{
		builder: builder,
	}
}

func (d *director) setBuilder(builder iBuilder) {
	d.builder = builder
}

func (d *director) buildHouse() house {
	d.builder.setNumFloor()
	d.builder.setDoorType()
	d.builder.setWindowType()
	return d.builder.getHouse()
}
