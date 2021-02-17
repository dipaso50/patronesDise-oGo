package house

import "patrones/command/commands"

type Interruptor struct {
	onCommand  commands.Command
	offCommand commands.Command
}

func NewInterruptor(onCmd, offcmd commands.Command) *Interruptor {
	return &Interruptor{
		onCommand:  onCmd,
		offCommand: offcmd,
	}
}

func (i *Interruptor) Encender() {
	i.onCommand.Execute()
}

func (i *Interruptor) Apagar() {
	i.offCommand.Execute()
}
