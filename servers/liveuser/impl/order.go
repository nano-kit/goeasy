package impl

import (
	"context"
	"time"

	"github.com/nano-kit/goeasy/servers/liveuser"
	"github.com/uptrace/bun"
)

type OrderService struct {
	sqlDB *bun.DB
}

type Order struct {
	ID       uint64     `bun:",pk"` // 订单编号
	UID      string     // 下单的用户ID
	State    OrderState `bun:"type:integer"` // 订单状态
	Amount   int32      // 商品总额（单位：分）
	Discount int32      // 折扣（单位：分）
	Pay      int32      // 实付款（单位：分）
	PayAt    time.Time  //  支付时间

	CreatedAt time.Time // 下单时间
	UpdatedAt time.Time // 订单状态更新时间
	DeletedAt time.Time
}

type OrderProduct struct {
	OrderID  uint64 // 订单编号
	ID       string // 商品编号
	Name     string // 商品名称
	Price    int32  // 购买时的商品价格（单位：分）
	Count    int32  // 购买的数量
	Snapshot uint64 // 购买时的商品快照
	Detail   string // 商品详情（JSON格式）
}

type OrderState struct {
	liveuser.OrderRecord_State
}

func (o *OrderService) Init(sqlDB *bun.DB) {
	o.sqlDB = sqlDB
}

func (o *OrderService) Create(ctx context.Context, req *liveuser.CreateOrderReq, res *liveuser.CreateOrderRes) error {
	return nil
}

func (o *OrderService) List(ctx context.Context, req *liveuser.ListOrderReq, res *liveuser.ListOrderRes) error {
	return nil
}
