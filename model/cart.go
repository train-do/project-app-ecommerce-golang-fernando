package model

type Cart struct {
	Id               int     `json:"id,omitempty"`
	UserId           int     `json:"userId,omitempty"`
	ProductVariantId int     `json:"productVariantId,omitempty"`
	Name             string  `json:"name,omitempty"`
	Variant          string  `json:"variant,omitempty"`
	Price            float32 `json:"price,omitempty"`
	Qty              int     `json:"qty,omitempty"`
}
