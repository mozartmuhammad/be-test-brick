package entity

type ScrapingTokopediaProductListResponse struct {
	PDPUrl   string
	Name     string
	ImageUrl string
	Price    int64
	Rating   float64
	ShopName string
}

type Product struct {
	ID          string
	Marketplace string
	Name        string
	Description string
	ImageUrl    string
	Price       float64
	Rating      float64
	ShopName    string
}
