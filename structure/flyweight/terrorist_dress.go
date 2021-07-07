package flyweight

type terroristDress struct {
	color string
}

func (t *terroristDress) getColor() string {
	return t.color
}

func newTerroristDress(color string) *terroristDress {
	return &terroristDress{
		color: color,
	}
}
