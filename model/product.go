package model

import "time"

type Product struct {
	Id                 int
	CategoryId         int `json:"categoryId,omitempty"`
	Category           string
	ProductVariantId   []int
	Name               string
	Description        string
	Stock              []int
	Variant            []string
	Price              float32
	Discount           int
	PriceAfterDiscount float32
	Rating             float32
	QtySold            int
	IsBestSelling      bool
	IsNew              bool
	ImageUrl           []string `json:"imageUrl,omitempty"`
	CreatedAt          time.Time
}
type Category struct {
	Id   int
	Name string
}
type Banner struct {
	Id        int
	BannerUrl string
	Title     string
	Subtitle  string
	PathPage  string
	TimeStart string
	TimeEnd   string
}
type ShowcaseProduct struct {
	Id        int
	BannerUrl string
	Title     string
	Subtitle  string
	ProductId int
	TimeStart string
	TimeEnd   string
}
type QueryProduct struct {
	Page          int
	Name          string
	Category      string
	IsBestSelling bool
}
