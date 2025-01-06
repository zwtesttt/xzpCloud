package domain

type Order struct {
	id          string
	userId      string
	totalAmount float64
	status      OrderStatus
	items       []*Item
	createdAt   int64
	updatedAt   int64
	deletedAt   int64
}

type OrderStatus int

// 定义订单状态的常量
const (
	OrderStatusPending   OrderStatus = iota // 0: 待支付
	OrderStatusPaid                         // 1: 已支付
	OrderStatusCancelled                    // 2: 已取消
)

type Item struct {
	productId string
	quantity  int
	price     float64
}

func NewOrder(id, userId string, totalAmount float64, status OrderStatus, items []*Item, createdAt, updatedAt, deletedAt int64) *Order {
	return &Order{
		id:          id,
		userId:      userId,
		totalAmount: totalAmount,
		status:      status,
		items:       items,
		createdAt:   createdAt,
		updatedAt:   updatedAt,
		deletedAt:   deletedAt,
	}
}

func NewOrderItem(productId string, quantity int, price float64) *Item {
	return &Item{
		productId: productId,
		quantity:  quantity,
		price:     price,
	}
}

func (o *Order) Id() string {
	return o.id
}

func (o *Order) UserId() string {
	return o.userId
}

func (o *Order) TotalAmount() float64 {
	return o.totalAmount
}

func (o *Order) Status() OrderStatus {
	return o.status
}

func (o *Order) SetStatus(status OrderStatus) {
	o.status = status
}

func (o *Order) Items() []*Item {
	return o.items
}

func (o *Order) CreatedAt() int64 {
	return o.createdAt
}

func (o *Order) UpdatedAt() int64 {
	return o.updatedAt
}

func (o *Order) DeletedAt() int64 {
	return o.deletedAt
}

func (i *Item) ProductId() string {
	return i.productId
}

func (i *Item) Quantity() int {
	return i.quantity
}

func (i *Item) Price() float64 {
	return i.price
}
