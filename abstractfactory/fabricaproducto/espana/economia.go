package espana

import "patrones/abstractfactory/fabricaproducto"

type EconomiaEspaniola struct{}

type Jamon struct {
	pesoKg     float32
	precioKilo float32
}

func (j Jamon) GetPrice() float64 {
	return float64(j.pesoKg * j.precioKilo)
}

func (j Jamon) GetDescription() string {
	return "Jamon"
}

func (e EconomiaEspaniola) CreateProducto() fabricaproducto.Producto {
	return createProductoJamon()
}

func createProductoJamon() fabricaproducto.Producto {
	return Jamon{5, 13}
}
