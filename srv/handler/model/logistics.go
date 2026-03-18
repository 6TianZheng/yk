package model

import "time"

// Logistics 物流配送表
type Logistics struct {
	ID          int64     `json:"id"`           // 物流ID
	OrderID     int64     `json:"order_id"`     // 订单ID
	OrderNo     string    `json:"order_no"`     // 订单编号
	MemberID    int64     `json:"member_id"`    // 会员ID
	ShipperName string    `json:"shipper_name"` // 物流公司
	TrackingNo  string    `json:"tracking_no"`  // 物流单号
	Status      int       `json:"status"`       // 状态: 0-待发货, 1-已发货, 2-运输中, 3-已签收
	ShipTime    time.Time `json:"ship_time"`    // 发货时间
	CreatedAt   time.Time `json:"created_at"`   // 创建时间
	UpdatedAt   time.Time `json:"updated_at"`   // 更新时间
}

// LogisticsTrack 物流轨迹表
type LogisticsTrack struct {
	ID          int64     `json:"id"`           // ID
	LogisticsID int64     `json:"logistics_id"` // 物流ID
	Content     string    `json:"content"`      // 轨迹内容
	ActionTime  time.Time `json:"action_time"`  // 操作时间
	CreatedAt   time.Time `json:"created_at"`   // 创建时间
}

// Shipper 物流公司表
type Shipper struct {
	ID        int64     `json:"id"`         // 公司ID
	Name      string    `json:"name"`       // 公司名称
	Code      string    `json:"code"`       // 编码
	Phone     string    `json:"phone"`      // 电话
	Sort      int       `json:"sort"`       // 排序
	Status    int       `json:"status"`     // 状态
	CreatedAt time.Time `json:"created_at"` // 创建时间
}

// ShippingMethod 配送方式表
type ShippingMethod struct {
	ID           int64     `json:"id"`            // 方式ID
	Name         string    `json:"name"`          // 方式名称
	Code         string    `json:"code"`          // 编码
	FirstFee     float64   `json:"first_fee"`     // 首费
	ContinueFee  float64   `json:"continue_fee"`  // 续费
	FreeShipping float64   `json:"free_shipping"` // 免费额度
	Status       int       `json:"status"`        // 状态
	CreatedAt    time.Time `json:"created_at"`    // 创建时间
}

// DeliveryStaff 配送员表
type DeliveryStaff struct {
	ID        int64     `json:"id"`         // 配送员ID
	Name      string    `json:"name"`       // 姓名
	Mobile    string    `json:"mobile"`     // 电话
	Status    int       `json:"status"`     // 状态
	CreatedAt time.Time `json:"created_at"` // 创建时间
}
