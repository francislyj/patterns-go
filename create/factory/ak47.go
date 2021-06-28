package factory

type ak47 struct {
	gun
}

func newAk47() iGun {
	return &ak47{
		gun: gun{
			name: "ak47",
			power: 47,
		},
	}
}
