package coupon

/*
Mercado Libre está implementando un nuevo beneficio para los

usuarios que más usan la
plataforma con un

cupón de cierto monto gratis

que les permitirá comprar tantos ítems marcados como favoritos sea posible,

siempre que no excedan el monto del cupón.
Para esto se está analizando construir una API que dado una lista de item_id y el monto
total pueda devolver la lista de ítems que maximice el total gastado.
*/

type CouponRepository interface {
	UseCoupon(coupon Coupon) error
}

type Coupon struct {
	Amount float32 `json:"amout"`
}
