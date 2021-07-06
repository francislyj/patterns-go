package wrapper

type cheeseTopping struct {
	pizza pizza
}

func (c *cheeseTopping) getPrice() int {
	prize := c.pizza.getPrice()
	return prize + 10
}
