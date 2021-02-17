package commands

type offCommand struct {
	onOffObject OnOff
}

func NewOffCommand(lp OnOff) *offCommand {
	return &offCommand{onOffObject: lp}
}

func (oc offCommand) Execute() {
	oc.onOffObject.Apagar()
}
