package scraping

type ScrapingTokopediaProduct struct {
	Data TokopediaProductCategory `json:"data"`
}

type TokopediaProductCategory struct {
	CategoryProducts TokopediaProductCategoryData `json:"CategoryProducts"`
}

type TokopediaProductCategoryData struct {
	Data []Detail `json:"data"`
}

type Detail struct {
	Name     string  `json:"name"`
	Url      string  `json:"url"`
	Price    int64   `json:"priceInt"`
	ImageUrl string  `json:"imageUrl"`
	Rating   float64 `json:"rating"`
	Shop     struct {
		Name string `json:"name"`
	} `json:"shop"`
}

type ScrapingTokopediaProductDetailResponse struct {
	Description string
}
