package model

import "time"

// MemberLevel 会员等级表
type MemberLevel struct {
	ID         int64     `json:"id"`          // 等级ID
	Name       string    `json:"name"`        // 等级名称
	Level      int       `json:"level"`       // 等级数值
	MinPoints  int       `json:"min_points"`  // 最低积分
	Discount   float64   `json:"discount"`    // 折扣
	PointsRate float64   `json:"points_rate"` // 积分比例
	Status     int       `json:"status"`      // 状态
	CreatedAt  time.Time `json:"created_at"`  // 创建时间
}

// MemberAddress 会员收货地址表
type MemberAddress struct {
	ID        int64     `json:"id"`         // 地址ID
	MemberID  int64     `json:"member_id"`  // 会员ID
	Name      string    `json:"name"`       // 收货人
	Mobile    string    `json:"mobile"`     // 电话
	Province  string    `json:"province"`   // 省份
	City      string    `json:"city"`       // 城市
	District  string    `json:"district"`   // 区
	Address   string    `json:"address"`    // 详细地址
	IsDefault int       `json:"is_default"` // 是否默认
	CreatedAt time.Time `json:"created_at"` // 创建时间
}

// MemberCart 购物车表
type MemberCart struct {
	ID        int64     `json:"id"`         // ID
	MemberID  int64     `json:"member_id"`  // 会员ID
	ProductID int64     `json:"product_id"` // 商品ID
	Quantity  int       `json:"quantity"`   // 数量
	Checked   int       `json:"checked"`    // 是否选中
	CreatedAt time.Time `json:"created_at"` // 创建时间
}
