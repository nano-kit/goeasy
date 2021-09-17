package catalog

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/micro/go-micro/v2/auth"
	"github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2/util/snowflake"
	"github.com/nano-kit/goeasy/internal/ierr"
	ijson "github.com/nano-kit/goeasy/internal/json"
	"github.com/uptrace/bun"
)

type Catalog struct {
	sqlDB *bun.DB
}

type Prod struct {
	bun.BaseModel `bun:"products"`
	Snapshot      uint64 `bun:",pk"`

	ID    string
	Name  string
	Price int32

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
	Operator  string
}

func (c *Catalog) Init(sqlDB *bun.DB) {
	c.sqlDB = sqlDB

	// create table
	models := []interface{}{
		(*Prod)(nil),
	}
	for _, model := range models {
		if _, err := c.sqlDB.NewCreateTable().
			IfNotExists().
			Model(model).
			Exec(context.TODO()); err != nil {
			logger.Errorf("can not create table: %v", err)
		}
	}

	// create index
	c.sqlDB.NewCreateIndex().IfNotExists().
		Model((*Prod)(nil)).Index("product_id").Column("id").
		Exec(context.TODO())
}

func (c *Catalog) List(ctx context.Context, req *ListReq, res *ListRes) error {
	return nil
}

func (c *Catalog) Set(ctx context.Context, req *SetReq, res *SetRes) error {
	// 检查参数
	acc, ok := auth.AccountFromContext(ctx)
	if !ok {
		return ierr.BadRequest("no account")
	}
	if err := req.Product.validate(); err != nil {
		return ierr.BadRequest("invalid product: %v", err)
	}

	// 查询已经存在的产品
	prod, ok := c.loadLatestProduct(ctx, req.Product.Id)
	if !ok {
		return c.insertNewProduct(ctx, acc, req.Product)
	}

	// 更新产品
	return c.updateProduct(ctx, acc, prod, req.Product)
}

func (p *Product) validate() error {
	if p == nil {
		return fmt.Errorf("product is nil")
	}
	if p.Id == "" {
		return fmt.Errorf("no product ID")
	}
	if p.Name == "" {
		return fmt.Errorf("no product name")
	}
	if p.PriceCent <= 0 {
		return fmt.Errorf("product price must be greater than zero")
	}
	return nil
}

func (c *Catalog) loadLatestProduct(ctx context.Context, id string) (prod *Prod, ok bool) {
	p := &Prod{}
	err := c.sqlDB.NewSelect().Model(p).
		Where("id = ?", id).
		Order("snapshot DESC").
		Limit(1).
		Scan(ctx)
	if err == sql.ErrNoRows {
		logger.Debugf("Catalog.loadLatestProduct: not found")
		return nil, false
	}
	logger.Debugf("Catalog.loadLatestProduct %q: %v", p.ID, ijson.Stringify(p))
	return p, true
}

func (c *Catalog) insertNewProduct(ctx context.Context, acc *auth.Account, newProd *Product) error {
	now := time.Now()
	p := &Prod{
		Snapshot:  snowflake.ID(),
		ID:        newProd.Id,
		Name:      newProd.Name,
		Price:     newProd.PriceCent,
		CreatedAt: now,
		UpdatedAt: now,
		Operator:  acc.ID,
	}
	_, err := c.sqlDB.NewInsert().Model(p).
		Exec(ctx)
	if err != nil {
		return ierr.Storage("Catalog.insertNewProduct %q: %v", p.ID, err)
	}
	return nil
}

func (c *Catalog) updateProduct(ctx context.Context, acc *auth.Account, oldProd *Prod, newProd *Product) error {
	if oldProd.equals(newProd) {
		return nil
	}

	now := time.Now()
	p := oldProd
	p.Snapshot = snowflake.ID()
	p.ID = newProd.Id
	p.Name = newProd.Name
	p.Price = newProd.PriceCent
	p.CreatedAt = now
	p.UpdatedAt = now
	p.Operator = acc.ID
	_, err := c.sqlDB.NewInsert().Model(p).
		Exec(ctx)
	if err != nil {
		return ierr.Storage("Catalog.updateProduct %q: %v", p.ID, err)
	}
	return nil
}

func (p *Prod) equals(pp *Product) bool {
	if p.DeletedAt.IsZero() && // not deleted
		p.ID == pp.Id &&
		p.Name == pp.Name &&
		p.Price == pp.PriceCent {
		return true
	}
	return false
}
