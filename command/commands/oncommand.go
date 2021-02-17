package commands

type onCommand struct {
	onOffObject OnOff
}

func NewOnCommand(lp OnOff) *onCommand {
	return &onCommand{onOffObject: lp}
}

func (oc onCommand) Execute() {
	oc.onOffObject.Encender()
}
