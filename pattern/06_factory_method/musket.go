package factory_method

type musket struct {
	Gun
}

func newMusket() Weapon {
	return &musket{
		Gun: Gun{
			name:  "Musket gun",
			power: 1,
		},
	}
}
