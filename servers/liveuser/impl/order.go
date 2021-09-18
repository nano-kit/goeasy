package impl

import (
	"context"
	"fmt"
	"math"
	"time"

	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/auth"
	"github.com/micro/go-micro/v2/client"
	"github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2/util/snowflake"
	"github.com/nano-kit/goeasy/internal/ierr"
	"github.com/nano-kit/goeasy/internal/ijson"
	"github.com/nano-kit/goeasy/internal/itime"
	"github.com/nano-kit/goeasy/servers/catalog"
	"github.com/nano-kit/goeasy/servers/liveuser"
	"github.com/uptrace/bun"
)

type OrderService struct {
	sqlDB          *bun.DB
	microClient    client.Client
	catelogService catalog.CatalogService
}

type Order struct {
	ID       uint64     `bun:",pk"` // 订单编号
	UID      string     // 下单的用户ID
	State    OrderState // 订单状态
	Amount   int32      // 商品总额（单位：分）
	Discount int32      // 折扣（单位：分）
	Pay      int32      // 实付款（单位：分）
	PayAt    time.Time  //  支付时间

	CreatedAt time.Time // 下单时间
	UpdatedAt time.Time // 订单状态更新时间
	DeletedAt time.Time
}

type OrderProduct struct {
	OrderID   uint64 // 订单编号
	ProductID string // 商品编号
	Name      string // 商品名称
	Price     int32  // 购买时的商品价格（单位：分）
	Count     int32  // 购买的数量
	Snapshot  uint64 // 购买时的商品快照
	Detail    string // 商品详情（JSON格式）
}

type OrderState liveuser.OrderRecord_State

func (o *Order) AfterCreateTable(ctx context.Context, query *bun.CreateTableQuery) error {
	_, err := query.DB().NewCreateIndex().
		IfNotExists().
		Model((*Order)(nil)).
		Index("order_uid").
		Column("uid").
		Exec(ctx)
	return err
}

func (p *OrderProduct) AfterCreateTable(ctx context.Context, query *bun.CreateTableQuery) error {
	_, err := query.DB().NewCreateIndex().
		IfNotExists().
		Model((*OrderProduct)(nil)).
		Index("order_product_order_id").
		Column("order_id").
		Exec(ctx)
	return err
}

func (o *OrderService) Init(sqlDB *bun.DB) {
	o.sqlDB = sqlDB

	// create table
	models := []interface{}{
		(*Order)(nil),
		(*OrderProduct)(nil),
	}
	for _, model := range models {
		if _, err := o.sqlDB.NewCreateTable().
			IfNotExists().
			Model(model).
			Exec(context.TODO()); err != nil {
			logger.Errorf("can not create table: %v", err)
		}
	}
}

func (o *OrderService) InitMicroClient(serivce micro.Service, namespace string) {
	o.microClient = &clientWrapper{
		Client:    serivce.Client(),
		namespace: namespace,
	}
	o.catelogService = catalog.NewCatalogService(namespace+".service.catalog", o.microClient)
}

func (o *OrderService) Create(ctx context.Context, req *liveuser.CreateOrderReq, res *liveuser.CreateOrderRes) error {
	// 检查参数
	acc, ok := auth.AccountFromContext(ctx)
	if !ok {
		return ierr.BadRequest("no account")
	}
	if err := o.validateCreateRequest(req); err != nil {
		return ierr.BadRequest("invalid request: %v", err)
	}

	// 查询商品
	snapshots := make([]uint64, len(req.Products))
	for i, p := range req.Products {
		snapshots[i] = p.ProductSnapshot
	}
	snapshotsResponse, err := o.catelogService.FindBySnapshot(ctx, &catalog.FindBySnapshotReq{
		Snapshots: snapshots,
	})
	if err != nil {
		return ierr.Internal("catelogService.FindBySnapshot: %v", err)
	}
	products := snapshotsResponse.Products
	if len(products) == 0 {
		return ierr.BadRequest("catelogService.FindBySnapshot: no product")
	}

	// 辅助函数
	countOfCatalogProduct := func(p *catalog.Product, order []*liveuser.OrderProduct) (count int32) {
		for _, op := range order {
			if p.Snapshot == op.ProductSnapshot {
				return op.Count
			}
		}
		return 0
	}

	// 创建订单
	orderID := snowflake.ID()
	orderProducts := make([]*OrderProduct, len(products))
	var amount int32
	for i, p := range products {
		op := &OrderProduct{
			OrderID:   orderID,
			ProductID: p.Id,
			Name:      p.Name,
			Price:     p.PriceCent,
			Count:     countOfCatalogProduct(p, req.Products),
			Snapshot:  p.Snapshot,
			Detail:    ijson.Stringify(p),
		}
		amount += op.Price * op.Count
		orderProducts[i] = op
	}
	order := &Order{
		ID:        orderID,
		UID:       acc.ID,
		State:     OrderState(liveuser.OrderRecord_CREATED),
		Amount:    amount,
		Pay:       amount,
		CreatedAt: time.Now(),
	}

	// 存储订单
	if err := o.saveOrder(ctx, order, orderProducts); err != nil {
		return ierr.Storage("CreateOrder: %v", err)
	}
	res.Order = serializeOrder(order, orderProducts)
	return nil
}

func (o *OrderService) List(ctx context.Context, req *liveuser.ListOrderReq, res *liveuser.ListOrderRes) error {
	// 检查参数
	acc, ok := auth.AccountFromContext(ctx)
	if !ok {
		return ierr.BadRequest("no account")
	}
	if req.Cursor == 0 {
		req.Cursor = math.MaxUint64
	}

	// 获取所有订单
	var orders []*Order
	err := o.sqlDB.NewSelect().Model(&orders).
		Where("id < ?", req.Cursor).
		Where("uid = ?", acc.ID).
		Order("id DESC").
		Limit(10).
		Scan(ctx)
	if err != nil {
		return ierr.Storage("ListOrder: %v", err)
	}
	if len(orders) == 0 {
		return nil
	}

	// 获取所有订单的ID
	orderIDs := make([]uint64, len(orders))
	for i, o := range orders {
		orderIDs[i] = o.ID
	}
	res.Cursor = orderIDs[len(orders)-1]

	// 获取订单内的商品
	var orderProducts []*OrderProduct
	err = o.sqlDB.NewSelect().Model(&orderProducts).
		Where("order_id IN (?)", bun.In(orderIDs)).
		Order("product_id").
		Scan(ctx)
	if err != nil {
		return ierr.Storage("ListOrderProducts: %v", err)
	}

	// 组合订单和商品
	res.Orders = make([]*liveuser.OrderRecord, len(orders))
	for i, thisOrder := range orders {
		var thisOrderProducts []*OrderProduct
		for _, orderProduct := range orderProducts {
			if orderProduct.OrderID == thisOrder.ID {
				thisOrderProducts = append(thisOrderProducts, orderProduct)
			}
		}
		res.Orders[i] = serializeOrder(thisOrder, thisOrderProducts)
	}
	return nil
}

func (o *OrderService) validateCreateRequest(req *liveuser.CreateOrderReq) error {
	if len(req.Products) == 0 {
		return fmt.Errorf("no product")
	}
	for _, p := range req.Products {
		if p.ProductSnapshot == 0 {
			return fmt.Errorf("no product snapshot")
		}
		if p.Count == 0 {
			return fmt.Errorf("product count is zero")
		}
	}
	return nil
}

func (o *OrderService) saveOrder(ctx context.Context, order *Order, orderProducts []*OrderProduct) error {
	err := o.sqlDB.RunInTx(ctx, nil, func(ctx context.Context, tx bun.Tx) (err error) {
		_, err = tx.NewInsert().Model(&orderProducts).Exec(ctx)
		if err != nil {
			return err
		}
		_, err = tx.NewInsert().Model(order).Exec(ctx)
		return err
	})
	return err
}

func serializeOrder(order *Order, orderProducts []*OrderProduct) *liveuser.OrderRecord {
	r := &liveuser.OrderRecord{
		Id:        order.ID,
		Uid:       order.UID,
		State:     liveuser.OrderRecord_State(order.State),
		Amount:    order.Amount,
		Discount:  order.Discount,
		Pay:       order.Pay,
		PayAt:     itime.MakeTimestamp(order.PayAt),
		Products:  make([]*liveuser.OrderProduct, len(orderProducts)),
		CreatedAt: itime.MakeTimestamp(order.CreatedAt),
		UpdatedAt: itime.MakeTimestamp(order.UpdatedAt),
		DeletedAt: itime.MakeTimestamp(order.DeletedAt),
	}
	for i, op := range orderProducts {
		r.Products[i] = &liveuser.OrderProduct{
			OrderId:         op.OrderID,
			ProductId:       op.ProductID,
			Name:            op.Name,
			Price:           op.Price,
			Count:           op.Count,
			ProductSnapshot: op.Snapshot,
			Detail:          op.Detail,
		}
	}
	return r
}
