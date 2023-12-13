package product

const (
	InsertProductQuery = `INSERT INTO product (id, marketplace, name, description, image_url, price, rating, shop_name)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`
)
