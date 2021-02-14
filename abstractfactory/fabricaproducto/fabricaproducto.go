package fabricaproducto

type Producto interface {
	GetDescription() string
	GetPrice() float64
}
type ProductoFabrica interface {
	CreateProducto() Producto
}
