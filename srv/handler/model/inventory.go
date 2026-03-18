package model

import "time"

// Inventory 库存表
type Inventory struct {
	ID        int64     `json:"id"`         // ID
	ProductID int64     `json:"product_id"` // 商品ID
	SKUID     int64     `json:"sku_id"`     // SKU ID
	Stock     int       `json:"stock"`      // 库存
	LockStock int       `json:"lock_stock"` // 锁定库存
	Status    int       `json:"status"`     // 状态
	CreatedAt time.Time `json:"created_at"` // 创建时间
	UpdatedAt time.Time `json:"updated_at"` // 更新时间
}

// Warehouse 仓库表
type Warehouse struct {
	ID        int64     `json:"id"`         // 仓库ID
	Name      string    `json:"name"`       // 仓库名称
	Code      string    `json:"code"`       // 编码
	Address   string    `json:"address"`    // 地址
	Manager   string    `json:"manager"`    // 负责人
	Mobile    string    `json:"mobile"`     // 电话
	Status    int       `json:"status"`     // 状态
	CreatedAt time.Time `json:"created_at"` // 创建时间
}

// ProductSKU 商品SKU表
type ProductSKU struct {
	ID         int64     `json:"id"`          // SKU ID
	ProductID  int64     `json:"product_id"`  // 商品ID
	SKUCode    string    `json:"sku_code"`    // SKU编码
	AttrValues string    `json:"attr_values"` // 属性值
	Price      float64   `json:"price"`       // 价格
	Stock      int       `json:"stock"`       // 库存
	Status     int       `json:"status"`      // 状态
	CreatedAt  time.Time `json:"created_at"`  // 创建时间
}

// InventoryLog 库存日志表
type InventoryLog struct {
	ID          int64     `json:"id"`           // ID
	InventoryID int64     `json:"inventory_id"` // 库存ID
	ProductID   int64     `json:"product_id"`   // 商品ID
	Type        int       `json:"type"`         // 类型: 1-入库, 2-出库
	ChangeStock int       `json:"change_stock"` // 变更数量
	OrderNo     string    `json:"order_no"`     // 订单号
	Remark      string    `json:"remark"`       // 备注
	CreatedAt   time.Time `json:"created_at"`   // 创建时间
}
