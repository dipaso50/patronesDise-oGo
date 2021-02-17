package commands

type Command interface {
	Execute()
}

type OnOff interface {
	Encender()
	Apagar()
}
