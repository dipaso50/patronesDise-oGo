package house

import "fmt"

type Lampara struct {
}

func (lp Lampara) Encender() {
	fmt.Println("Lampara encendida")
}

func (lp Lampara) Apagar() {
	fmt.Println("Lampara apagada")
}
