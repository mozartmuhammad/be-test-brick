package scraping

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/gocolly/colly"
	"github.com/mozartmuhammad/be-test-brick/src/entity"
)

func (d *domain) ScrapingTokopediaProductList(ctx context.Context, page int, pageSize int) ([]entity.ScrapingTokopediaProductListResponse, error) {
	var result []entity.ScrapingTokopediaProductListResponse

	client := &http.Client{}
	param := fmt.Sprintf(`[{"operationName":"SearchProductQuery","variables":{"params":"page=1&ob=&identifier=handphone-tablet_handphone&sc=24&user_id=0&rows=%d&start=1&source=directory&device=desktop&page=%d&related=true&st=product&safe_search=false","adParams":"page=1&page=1&dep_id=24&ob=&ep=product&item=15&src=directory&device=desktop&user_id=0&minimum_item=15&start=1&no_autofill_range=5-14"},"query":"query SearchProductQuery($params: String, $adParams: String) {\n  CategoryProducts: searchProduct(params: $params) {\n    count\n    data: products {\n      id\n      url\n      imageUrl: image_url\n      imageUrlLarge: image_url_700\n      catId: category_id\n      gaKey: ga_key\n      countReview: count_review\n      discountPercentage: discount_percentage\n      preorder: is_preorder\n      name\n      price\n      priceInt: price_int\n      original_price\n      rating\n      wishlist\n      labels {\n        title\n        color\n        __typename\n      }\n      badges {\n        imageUrl: image_url\n        show\n        __typename\n      }\n      shop {\n        id\n        url\n        name\n        goldmerchant: is_power_badge\n        official: is_official\n        reputation\n        clover\n        location\n        __typename\n      }\n      labelGroups: label_groups {\n        position\n        title\n        type\n        __typename\n      }\n      __typename\n    }\n    __typename\n  }\n  displayAdsV3(displayParams: $adParams) {\n    data {\n      id\n      ad_ref_key\n      redirect\n      sticker_id\n      sticker_image\n      productWishListUrl: product_wishlist_url\n      clickTrackUrl: product_click_url\n      shop_click_url\n      product {\n        id\n        name\n        wishlist\n        image {\n          imageUrl: s_ecs\n          trackerImageUrl: s_url\n          __typename\n        }\n        url: uri\n        relative_uri\n        price: price_format\n        campaign {\n          original_price\n          discountPercentage: discount_percentage\n          __typename\n        }\n        wholeSalePrice: wholesale_price {\n          quantityMin: quantity_min_format\n          quantityMax: quantity_max_format\n          price: price_format\n          __typename\n        }\n        count_talk_format\n        countReview: count_review_format\n        category {\n          id\n          __typename\n        }\n        preorder: product_preorder\n        product_wholesale\n        free_return\n        isNewProduct: product_new_label\n        cashback: product_cashback_rate\n        rating: product_rating\n        top_label\n        bottomLabel: bottom_label\n        __typename\n      }\n      shop {\n        image_product {\n          image_url\n          __typename\n        }\n        id\n        name\n        domain\n        location\n        city\n        tagline\n        goldmerchant: gold_shop\n        gold_shop_badge\n        official: shop_is_official\n        lucky_shop\n        uri\n        owner_id\n        is_owner\n        badges {\n          title\n          image_url\n          show\n          __typename\n        }\n        __typename\n      }\n      applinks\n      __typename\n    }\n    template {\n      isAd: is_ad\n      __typename\n    }\n    __typename\n  }\n}\n"}]`,
		pageSize, page)
	var data = strings.NewReader(param)
	req, err := http.NewRequest("POST", "https://gql.tokopedia.com/graphql/SearchProductQuery", data)
	if err != nil {
		return result, err
	}
	req.Header.Set("sec-ch-ua", `"Google Chrome";v="119", "Chromium";v="119", "Not?A_Brand";v="24"`)
	req.Header.Set("Tkpd-UserId", "0")
	req.Header.Set("X-Version", "bbe393e")
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/119.0.0.0 Safari/537.36")
	req.Header.Set("iris_session_id", "d3d3LnRva29wZWRpYS5jb20=.8d2a43e50e2ee490998f8493501ad95b.1702363359430")
	req.Header.Set("content-type", "application/json")
	req.Header.Set("accept", "*/*")
	req.Header.Set("Referer", "https://www.tokopedia.com/p/handphone-tablet/handphone")
	req.Header.Set("X-Source", "tokopedia-lite")
	req.Header.Set("x-device", "desktop-0.0")
	req.Header.Set("X-Tkpd-Lite-Service", "zeus")
	req.Header.Set("sec-ch-ua-platform", `"macOS"`)
	resp, err := client.Do(req)
	if err != nil {
		return result, err
	}
	defer resp.Body.Close()
	raw, err := io.ReadAll(resp.Body)
	if err != nil {
		return result, err
	}

	var list []ScrapingTokopediaProduct
	_ = json.Unmarshal(raw, &list)
	for _, item := range list {
		for _, product := range item.Data.CategoryProducts.Data {
			result = append(result, entity.ScrapingTokopediaProductListResponse{
				PDPUrl:   product.Url,
				Name:     product.Name,
				ImageUrl: product.ImageUrl,
				Price:    product.Price,
				Rating:   product.Rating,
				ShopName: product.Shop.Name,
			})
		}
	}

	return result, nil
}

func (d *domain) ScrapingTokopediaProductDetail(ctx context.Context, pdpUrl string) ScrapingTokopediaProductDetailResponse {
	var resp ScrapingTokopediaProductDetailResponse

	c := colly.NewCollector(
		// turning on the asynchronous request mode in Colly
		colly.Async(true),
	)

	// setting a valid User-Agent header
	c.UserAgent = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/119.0.0.0 Safari/537.36"

	c.OnHTML("[data-testid=lblPDPDescriptionProduk]", func(e *colly.HTMLElement) {
		resp.Description = e.Text
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visit", r.URL.String())
	})

	c.Visit(pdpUrl)
	c.Wait()

	return resp
}
