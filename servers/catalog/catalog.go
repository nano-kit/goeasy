package catalog

import (
	"context"

	"github.com/uptrace/bun"
)

type Catalog struct {
	sqlDB *bun.DB
}

func (c *Catalog) Init(sqlDB *bun.DB) {
	c.sqlDB = sqlDB
}

func (c *Catalog) List(ctx context.Context, req *ListReq, res *ListRes) error {
	return nil
}
