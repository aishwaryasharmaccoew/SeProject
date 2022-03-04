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

// ProductList contains a list of Products
type ProductList struct {
	// A list of Products
	Items []*ProductSQL `json:"items"`
	// The id to query the next page
	NextPageID string `json:"next_page_id,omitempty" example:"10"`
}
