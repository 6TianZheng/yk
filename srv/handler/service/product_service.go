package service

import (
	"context"
	"errors"
	"fmt"
	"lx/srv/basic/config"
	__ "lx/srv/basic/proto"
	"lx/srv/handler/model"
	"lx/srv/pkg"
)

type Server struct {
	__.UnimplementedProductServiceServer
}

// CreateProduct 创建商品
func (s *Server) CreateProduct(_ context.Context, req *__.CreateProductRequest) (*__.Product, error) {
	product := model.Product{
		CategoryID:  req.CategoryId,
		BrandID:     req.BrandId,
		Name:        req.Name,
		Images:      req.Images,
		Description: req.Description,
		Price:       req.Price,
		Stock:       int(req.Stock),
		Status:      int(req.Status),
	}

	if err := config.DB.Create(&product).Error; err != nil {
		return nil, fmt.Errorf("创建商品失败: %v", err)
	}

	return s.toProto(&product), nil
}

// UpdateProduct 更新商品
func (s *Server) UpdateProduct(_ context.Context, req *__.UpdateProductRequest) (*__.Product, error) {
	var product model.Product
	if err := config.DB.First(&product, req.Id).Error; err != nil {
		return nil, fmt.Errorf("商品不存在: %v", err)
	}

	updates := map[string]interface{}{
		"category_id": req.CategoryId,
		"brand_id":    req.BrandId,
		"name":        req.Name,
		"images":      req.Images,
		"description": req.Description,
		"price":       req.Price,
		"stock":       req.Stock,
		"status":      req.Status,
	}

	if err := config.DB.Model(&product).Updates(updates).Error; err != nil {
		return nil, fmt.Errorf("更新商品失败: %v", err)
	}

	return s.toProto(&product), nil
}

// DeleteProduct 删除商品
func (s *Server) DeleteProduct(_ context.Context, req *__.DeleteProductRequest) (*__.DeleteProductResponse, error) {
	if err := config.DB.Delete(&model.Product{}, req.Id).Error; err != nil {
		return &__.DeleteProductResponse{Success: false}, fmt.Errorf("删除商品失败: %v", err)
	}
	return &__.DeleteProductResponse{Success: true}, nil
}

// GetProduct 获取商品详情
func (s *Server) GetProduct(_ context.Context, req *__.GetProductRequest) (*__.Product, error) {
	var product model.Product
	if err := config.DB.First(&product, req.Id).Error; err != nil {
		return nil, fmt.Errorf("商品不存在: %v", err)
	}
	return s.toProto(&product), nil
}

// ListProduct 商品列表
func (s *Server) ListProduct(_ context.Context, req *__.ListProductRequest) (*__.ListProductResponse, error) {
	var products []model.Product
	query := config.DB.Model(&model.Product{})

	// 条件筛选
	if req.CategoryId > 0 {
		query = query.Where("category_id = ?", req.CategoryId)
	}
	if req.BrandId > 0 {
		query = query.Where("brand_id = ?", req.BrandId)
	}
	if req.Keyword != "" {
		query = query.Where("name LIKE ?", "%"+req.Keyword+"%")
	}
	if req.Status > 0 {
		query = query.Where("status = ?", req.Status)
	}

	// 统计总数
	var total int64
	query.Count(&total)

	// 分页
	page, pageSize := int(req.Page), int(req.PageSize)
	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}
	offset := (page - 1) * pageSize
	query = query.Offset(offset).Limit(pageSize)

	if err := query.Find(&products).Error; err != nil {
		return nil, fmt.Errorf("查询商品列表失败: %v", err)
	}

	// 转换
	protoProducts := make([]*__.Product, len(products))
	for i := range products {
		protoProducts[i] = s.toProto(&products[i])
	}

	return &__.ListProductResponse{
		Products: protoProducts,
		Total:    int32(total),
	}, nil
}

// toProto 转换为proto对象
func (s *Server) toProto(p *model.Product) *__.Product {
	return &__.Product{
		Id:          int64(p.ID),
		CategoryId:  p.CategoryID,
		BrandId:     p.BrandID,
		Name:        p.Name,
		Images:      p.Images,
		Description: p.Description,
		Price:       p.Price,
		Stock:       int32(p.Stock),
		Status:      int32(p.Status),
		CreatedAt:   p.CreatedAt.Unix(),
		UpdatedAt:   p.UpdatedAt.Unix(),
	}
}

// OrderItem 创建订单和订单明细
func (s *Server) OrderItem(_ context.Context, in *__.OrderItemAddRequest) (*__.OrderItemAddResponse, error) {
	orderSn := pkg.OrderSn()
	total := 0.0

	var list []*model.OrderItem

	for _, item := range in.List {
		var product model.Product
		err := config.DB.Where("id = ?", item.ProductId).First(&product).Error
		if err != nil {
			return nil, errors.New("商品不存在")
		}

		if product.Status != 1 {
			return nil, errors.New("商品未上架")
		}

		if item.Quantity > int64(product.Stock) {
			return nil, errors.New("库存不足")
		}

		total += product.Price * float64(item.Quantity)

		list = append(list, &model.OrderItem{
			ProductId:    int64(product.ID),
			ProductName:  product.Name,
			ProductPrice: product.Price,
			ProductImg:   product.Images,
			ProductBio:   product.Description,
			Quantity:     item.Quantity,
		})
	}

	order := model.Order{
		OrderSn:   orderSn,
		UserId:    in.UserId,
		PayType:   in.PayType,
		AddressId: in.AddressId,
		Status:    1,
	}

	err := config.DB.Create(&order).Error
	if err != nil {
		return nil, errors.New("订单创建失败")
	}

	for i := range list {
		list[i].OrderId = int64(order.ID)
	}

	err = config.DB.Create(&list).Error
	if err != nil {
		return nil, errors.New("订单明细创建失败")
	}

	payUrl := pkg.AliPay(orderSn, total)

	return &__.OrderItemAddResponse{
		OrderSn: orderSn,
		PayUrl:  payUrl,
		Total:   float32(total),
	}, nil
}
