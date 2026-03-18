package model

import (
	"gorm.io/gorm"
	"time"
)

// Product 商品表
type Product struct {
	gorm.Model
	CategoryID  int64   `json:"category_id"` // 分类ID
	BrandID     int64   `json:"brand_id"`    // 品牌ID
	Name        string  `json:"name"`        // 商品名称
	Images      string  `json:"images"`      // 商品图片
	Description string  `json:"description"` // 商品描述
	Price       float64 `json:"price"`       // 销售价格
	Stock       int     `json:"stock"`       // 库存
	Status      int     `json:"status"`      // 状态: 0-下架, 1-上架
}

func (p *Product) FindProductById(db *gorm.DB, id int64) error {
	return db.Debug().Where("id = ?", id).First(&p).Error
}

// ProductCategory 商品分类表
type ProductCategory struct {
	ID        int64     `json:"id"`         // 分类ID
	ParentID  int64     `json:"parent_id"`  // 父分类ID
	Name      string    `json:"name"`       // 分类名称
	Sort      int       `json:"sort"`       // 排序
	Status    int       `json:"status"`     // 状态
	CreatedAt time.Time `json:"created_at"` // 创建时间
}

// ProductBrand 品牌表
type ProductBrand struct {
	ID        int64     `json:"id"`         // 品牌ID
	Name      string    `json:"name"`       // 品牌名称
	Logo      string    `json:"logo"`       // Logo
	Sort      int       `json:"sort"`       // 排序
	Status    int       `json:"status"`     // 状态
	CreatedAt time.Time `json:"created_at"` // 创建时间
}
