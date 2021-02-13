package main

import (
	"patrones/decorator/linedecorator"
	"patrones/decorator/linewriter"
)

func main() {
	lwriter := linewriter.ConsoleWriter{}

	msg := "hi all!!!"

	lwriter.WriteLine(msg)

	//enriquecemos la instancia de ConsoleWriter con nueva funcionalidad
	upperWriter := linedecorator.NewUpperLineDecorator(&lwriter)
	historyWriter := linedecorator.NewHistoryLineDecorator(upperWriter)

	historyWriter.WriteLine(msg)
	historyWriter.WriteLine("adios")

	historyWriter.PrintHistory()
}
