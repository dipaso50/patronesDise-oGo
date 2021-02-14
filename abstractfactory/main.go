package main

import (
	"fmt"
	"patrones/abstractfactory/fabricaproducto"
	"patrones/abstractfactory/fabricaproducto/espana"
	"patrones/abstractfactory/fabricaproducto/mexico"
)

func main() {
	printData(mexico.EconomiaMexico{})
	printData(espana.EconomiaEspaniola{})
}

func printData(pr fabricaproducto.ProductoFabrica) {
	prod := pr.CreateProducto()
	fmt.Printf("Producto (%s) precio (%v euros) \n", prod.GetDescription(), prod.GetPrice())
}
