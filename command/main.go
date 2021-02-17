package main

import (
	"patrones/command/commands"
	"patrones/command/house"
)

func main() {

	lp := house.Lampara{}

	oncmd := commands.NewOnCommand(lp)
	offcmd := commands.NewOffCommand(lp)

	interruptor := house.NewInterruptor(oncmd, offcmd)

	interruptor.Apagar()
	interruptor.Encender()
	interruptor.Apagar()

}
