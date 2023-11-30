package command

type TurnOnCommand struct {
	device Device
}

func (c *TurnOnCommand) execute() {
	c.device.on()
}
