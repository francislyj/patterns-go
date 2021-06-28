package factory

type musket struct {
	gun
}

func newMusket() iGun {
	return &musket{
		gun: gun{power: 1,
			name: "musket"},
	}
}
