package product

import (
	"context"
	"database/sql"

	_ "github.com/lib/pq"
	"github.com/mozartmuhammad/be-test-brick/src/entity"
)

type DomainItf interface {
	UpsertProduct(ctx context.Context, v entity.Product) error
}

type domain struct {
	sql *sql.DB
}

type Options struct {
	Dsn string
}

func InitProduct(opts Options) *domain {
	sql, err := sql.Open("postgres", opts.Dsn)
	if err != nil {
		panic(err)
	}
	return &domain{
		sql: sql,
	}
}
