package model

type OrderProduct struct {
	CartId []int
}

type Order struct {
	Name          []string
	SubTotal      []float32
	TotalPrice    float32
	ShippingPrice int
}
