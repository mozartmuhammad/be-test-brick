package scraping

import (
	"context"

	"github.com/mozartmuhammad/be-test-brick/src/entity"
)

type DomainItf interface {
	ScrapingTokopediaProductList(ctx context.Context, page int, pageSize int) ([]entity.ScrapingTokopediaProductListResponse, error)
	ScrapingTokopediaProductDetail(ctx context.Context, pdpUrl string) ScrapingTokopediaProductDetailResponse
}

type domain struct {
}

func InitScraping() *domain {
	return &domain{}
}
