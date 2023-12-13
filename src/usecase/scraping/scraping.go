package scraping

import (
	"context"

	"github.com/mozartmuhammad/be-test-brick/src/domain/product"
	"github.com/mozartmuhammad/be-test-brick/src/domain/scraping"
)

type UsecaseItf interface {
	ScrapingTokopedia(ctx context.Context) error
}

type usecase struct {
	opt         Options
	scrapingDOM scraping.DomainItf
	productDOM  product.DomainItf
}

type Options struct {
	MaxGoRoutine int64
}

func Init(
	opt Options,
	scrapingDOM scraping.DomainItf,
	productDOM product.DomainItf,
) UsecaseItf {
	return &usecase{
		opt:         opt,
		scrapingDOM: scrapingDOM,
		productDOM:  productDOM,
	}
}
