package command

type TurnOffCommand struct {
	device Device
}

func (c *TurnOffCommand) execute() {
	c.device.off()
}
