package main

import (
	"context"
	"os"

	"github.com/mozartmuhammad/be-test-brick/src/domain"
	"github.com/mozartmuhammad/be-test-brick/src/domain/product"
	"github.com/mozartmuhammad/be-test-brick/src/usecase"
	"github.com/mozartmuhammad/be-test-brick/src/usecase/scraping"
)

var dom *domain.Domain

func main() {
	dbDsn := os.Getenv("DATABASE_URL")

	dom = domain.Init(domain.Options{
		Product: product.Options{
			Dsn: dbDsn,
		},
	})

	uc := usecase.Init(usecase.Options{
		Scraping: scraping.Options{},
	}, dom)

	uc.Scraping.ScrapingTokopedia(context.Background())
}
