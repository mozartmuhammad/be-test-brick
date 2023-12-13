package scraping

import (
	"context"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/google/uuid"
	"github.com/mozartmuhammad/be-test-brick/src/entity"
)

func (uc *usecase) ScrapingTokopedia(ctx context.Context) error {
	var (
		productList []entity.ScrapingTokopediaProductListResponse
		fName       = "result.csv"
		page        = 1
		pageSize    = 50
	)

	// fetch product list
	for ; page <= 2; page++ {
		list, _ := uc.scrapingDOM.ScrapingTokopediaProductList(ctx, page, pageSize)
		productList = append(productList, list...)
	}

	ch := make(chan entity.Product, 5)
	var wg sync.WaitGroup
	for i := range productList {
		wg.Add(1)
		go func(v entity.ScrapingTokopediaProductListResponse) {
			defer wg.Done()
			detail := uc.scrapingDOM.ScrapingTokopediaProductDetail(ctx, v.PDPUrl)

			ch <- entity.Product{
				ID:          uuid.New().String(),
				Marketplace: entity.MARKETPLACE_TOKOPEDIA,
				Name:        v.Name,
				Description: detail.Description,
				ImageUrl:    v.ImageUrl,
				Price:       float64(v.Price),
				Rating:      v.Rating,
				ShopName:    v.ShopName,
			}

		}(productList[i])
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	// write csv
	file, err := os.Create(fName)
	if err != nil {
		log.Fatalf("Could not create file, err: %q", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()
	for c := range ch {
		err = uc.productDOM.UpsertProduct(ctx, c)
		if err != nil {
			println(err.Error())
		}
		writer.Write([]string{c.Name, c.Description, c.Description, c.ImageUrl, fmt.Sprint(c.Price), fmt.Sprint(c.Rating), c.ShopName})
	}

	fmt.Println("Scraping tokopedia done.")
	return nil
}
