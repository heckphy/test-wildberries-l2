package factory_method

type Ak47 struct {
	Gun
}

func newAk47() Weapon {
	return &Ak47{
		Gun: Gun{
			name:  "AK47 gun",
			power: 4,
		},
	}
}
