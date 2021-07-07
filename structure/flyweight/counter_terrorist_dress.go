package flyweight

type counterTerroristDress struct {
	color string
}

func (c *counterTerroristDress) getColor() string {
	return c.color
}

func newCounterTerroristDress(color string) *counterTerroristDress {
	return &counterTerroristDress{
		color: color,
	}
}
