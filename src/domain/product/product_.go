package product

import (
	"context"

	"github.com/mozartmuhammad/be-test-brick/src/entity"
)

func (d *domain) UpsertProduct(ctx context.Context, v entity.Product) error {
	tx, err := d.sql.Begin()
	if err != nil {
		return err
	}

	query, err := tx.PrepareContext(ctx, InsertProductQuery)
	if err != nil {
		return err
	}
	defer query.Close()

	_, err = query.ExecContext(ctx,
		v.ID,
		v.Marketplace,
		v.Name,
		v.Description,
		v.ImageUrl,
		v.Price,
		v.Rating,
		v.ShopName,
	)
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}
