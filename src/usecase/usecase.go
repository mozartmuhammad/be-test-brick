package usecase

import (
	"github.com/mozartmuhammad/be-test-brick/src/domain"
	"github.com/mozartmuhammad/be-test-brick/src/usecase/scraping"
)

type Usecase struct {
	Scraping scraping.UsecaseItf
}

type Options struct {
	Scraping scraping.Options
}

func Init(
	conf Options,
	dom *domain.Domain,
) *Usecase {
	scrapingUC := scraping.Init(
		conf.Scraping, dom.Scraping, dom.Product,
	)

	return &Usecase{
		Scraping: scrapingUC,
	}
}
