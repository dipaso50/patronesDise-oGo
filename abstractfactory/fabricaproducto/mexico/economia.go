package mexico

import "patrones/abstractfactory/fabricaproducto"

type EconomiaMexico struct{}

type Mariachis struct {
	precioPorHora float32
}

func (j Mariachis) GetPrice() float64 {
	return float64(j.precioPorHora)
}

func (j Mariachis) GetDescription() string {
	return "Marichi las mañanitas"
}

func (e EconomiaMexico) CreateProducto() fabricaproducto.Producto {
	return contrataMariachis()
}

func contrataMariachis() fabricaproducto.Producto {
	return Mariachis{25}
}
