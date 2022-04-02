package models

type Product struct {
	Id           string
	Name         string
	Price        float32
	Category     string
	ImageUrl     string
	Product_url  string
	Colors       []string
	Key_features []string
	Tools        []string
}

type ProductSQL struct {
	Id           string
	Name         string
	Price        float32
	Category     string
	ImageUrl     string
	Product_url  string
	Colors       string
	Key_features string
	Tools        string
}
