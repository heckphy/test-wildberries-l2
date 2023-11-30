package factory_method

import "fmt"

func NewWeapon(weaponType string) (Weapon, error) {
	switch weaponType {
	case "ak47":
		return newAk47(), nil
	case "musket":
		return newMusket(), nil
	}

	return nil, fmt.Errorf("Wrong weapon type passed")
}
