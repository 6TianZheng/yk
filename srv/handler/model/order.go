package model

import (
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	OrderSn   string `gorm:"type:varchar(50);comment:'订单号'"`
	UserId    int64  `gorm:"type:int(11);comment:'用户id'"`
	PayType   int64  `gorm:"type:int(11);comment:'支付方式'"` //1支付宝
	AddressId int64  `gorm:"type:int(11);comment:'地址'"`
	Status    int64  `gorm:"type:int(11);comment:'订单状态'"`
}

func (o *Order) OrderAdd(db *gorm.DB) error {
	return db.Debug().Create(&o).Error
}

func (o *Order) OrderItemAdd(db *gorm.DB, list []*OrderItem) error {
	return db.Debug().Create(&list).Error
}

type OrderItem struct {
	gorm.Model
	OrderId      int64   `gorm:"type:int(11);comment:'订单id'"`
	ProductId    int64   `gorm:"type:int(11);comment:'菜品id'"`
	ProductName  string  `gorm:"type:varchar(50);comment:'菜品名称'"`
	ProductPrice float64 `gorm:"type:decimal(10,2);comment:'菜品单价'"`
	ProductImg   string  `gorm:"type:varchar(50);comment:'菜品图片'"`
	ProductBio   string  `gorm:"type:varchar(50);comment:'菜品详细'"`
	Quantity     int64   `gorm:"type:int(11);comment:'购买数量'"`
}

type Address struct {
	gorm.Model
	Name string `gorm:"type:varchar(50);comment:'收货人姓名'"`
	Tel  string `gorm:"type:varchar(11);comment:'手机号'"`
	Info string `gorm:"type:varchar(225);comment:'详细地址'"`
}
