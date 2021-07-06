package wrapper

type tomatoTopping struct {
	pizza pizza
}

func (t *tomatoTopping) getPrice() int {
	prize := t.pizza.getPrice()
	return prize + 8
}

