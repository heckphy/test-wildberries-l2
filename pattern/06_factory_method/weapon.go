package factory_method

type Weapon interface {
	setName(name string)
	setPower(power int)
	getName() string
	getPower() int
}
