package factory_abs

type nike struct {

}

func (n *nike) makeShoe() iShoe {
	return &nikeShoe{
		shoe{
			logo: "nike",
			size: 20,
		},
	}
}

func (n *nike) makeShirt() iShirt {
	return &nikeShirt{
		shirt{
			logo: "nike",
			size: 45,
		},
	}
}
