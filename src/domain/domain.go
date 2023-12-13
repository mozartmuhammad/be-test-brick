package domain

import (
	"github.com/mozartmuhammad/be-test-brick/src/domain/product"
	"github.com/mozartmuhammad/be-test-brick/src/domain/scraping"
)

type Domain struct {
	Product  product.DomainItf
	Scraping scraping.DomainItf
}

type Options struct {
	Product product.Options
}

func Init(
	conf Options,
) *Domain {
	productDom := product.InitProduct(
		conf.Product,
	)
	scrapingDom := scraping.InitScraping()

	return &Domain{
		Product:  productDom,
		Scraping: scrapingDom,
	}
}
