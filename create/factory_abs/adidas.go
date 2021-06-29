package factory_abs

type adidas struct {

}

func (a *adidas) makeShoe() iShoe {
	return &adidasShoe{
		shoe{
			size: 20,
			logo: "adidas",
		},
	}
}

func (a *adidas) makeShirt() iShirt {
	return &adidasShirt{
		shirt{
			logo: "adidas",
			size: 80,
		},
	}
}
