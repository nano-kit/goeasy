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
	"github.com/nano-kit/goeasy/internal/ijson"
	"github.com/nano-kit/goeasy/internal/itime"
	"github.com/uptrace/bun"
)

var (
	zeroTime = time.Time{}
)

type Catalog struct {
	sqlDB *bun.DB
}

type Prod struct {
	bun.BaseModel `bun:"products,alias:p"`
	Snapshot      uint64 `bun:",pk"`

	ID    string
	Name  string
	Price int32

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
	Operator  string
}

func (p *Prod) AfterCreateTable(ctx context.Context, query *bun.CreateTableQuery) error {
	_, err := query.DB().NewCreateIndex().
		IfNotExists().
		Model((*Prod)(nil)).
		Index("product_id").
		Column("id").
		Exec(ctx)
	return err
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
}

func (c *Catalog) List(ctx context.Context, req *ListReq, res *ListRes) error {
	// 检查参数
	_, ok := auth.AccountFromContext(ctx)
	if !ok {
		return ierr.BadRequest("no account")
	}

	products, err := c.load(ctx, nil)
	if err != nil {
		return ierr.Storage("Catalog.List: %v", err)
	}

	res.Products = make([]*Product, len(products))
	for i, p := range products {
		res.Products[i] = p.serialize()
	}
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

func (c *Catalog) Delete(ctx context.Context, req *DeleteReq, res *DeleteRes) error {
	// 检查参数
	acc, ok := auth.AccountFromContext(ctx)
	if !ok {
		return ierr.BadRequest("no account")
	}
	if req.ProductId == "" {
		return ierr.BadRequest("empty product identity")
	}

	// 查询已经存在的产品
	prod, ok := c.loadLatestProduct(ctx, req.ProductId)
	if !ok {
		return nil
	}

	// 删除产品
	return c.deleteProduct(ctx, acc, prod)
}

func (c *Catalog) FindByID(ctx context.Context, req *FindByIDReq, res *FindByIDRes) error {
	// 检查参数
	_, ok := auth.AccountFromContext(ctx)
	if !ok {
		return ierr.BadRequest("no account")
	}
	if len(req.ProductIds) == 0 {
		return ierr.BadRequest("no product")
	}

	products, err := c.load(ctx, req.ProductIds)
	if err != nil {
		return ierr.Storage("Catalog.FindByID: %v", err)
	}

	res.Products = make([]*Product, len(products))
	for i, p := range products {
		res.Products[i] = p.serialize()
	}
	return nil
}

func (c *Catalog) load(ctx context.Context, ids []string) ([]*Prod, error) {
	maxSnapshot := c.sqlDB.NewSelect().Model((*Prod)(nil)).
		Column("id").
		ColumnExpr("max(snapshot) AS snapshot").
		Group("id")

	if len(ids) > 0 {
		maxSnapshot.Where("id IN (?)", bun.In(ids))
	}

	var products []*Prod
	err := c.sqlDB.NewSelect().With("pm", maxSnapshot).
		Model(&products).
		Join("INNER JOIN pm").
		JoinOn("p.snapshot = pm.snapshot").
		JoinOn("p.id = pm.id").
		//Where("p.deleted_at = ?", zeroTime).
		Order("p.id").
		Scan(ctx)
	return products, err
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
	if !p.DeletedAt.IsZero() {
		logger.Debugf("Catalog.loadLatestProduct %q: deleted at %v", p.ID, p.DeletedAt)
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
	p.Name = newProd.Name
	p.Price = newProd.PriceCent
	p.UpdatedAt = now
	p.DeletedAt = zeroTime
	p.Operator = acc.ID
	_, err := c.sqlDB.NewInsert().Model(p).
		Exec(ctx)
	if err != nil {
		return ierr.Storage("Catalog.updateProduct %q: %v", p.ID, err)
	}
	return nil
}

func (c *Catalog) deleteProduct(ctx context.Context, acc *auth.Account, p *Prod) error {
	now := time.Now()
	p.Snapshot = snowflake.ID()
	p.UpdatedAt = now
	p.DeletedAt = now
	p.Operator = acc.ID
	_, err := c.sqlDB.NewInsert().Model(p).
		Exec(ctx)
	if err != nil {
		return ierr.Storage("Catalog.deleteProduct %q: %v", p.ID, err)
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

func (p *Prod) serialize() *Product {
	return &Product{
		Id:        p.ID,
		Name:      p.Name,
		PriceCent: p.Price,
		Snapshot:  p.Snapshot,
		CreatedAt: itime.MakeTimestamp(p.CreatedAt),
		UpdatedAt: itime.MakeTimestamp(p.UpdatedAt),
		DeletedAt: itime.MakeTimestamp(p.DeletedAt),
		Operator:  p.Operator,
	}
}
