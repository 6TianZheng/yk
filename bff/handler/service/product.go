package service

import (
	"context"
	"fmt"
	config2 "lx/srv/basic/config"
	"net/http"
	"strconv"

	"lx/bff/basic/config"
	"lx/bff/handler/request"
	"lx/srv/basic/proto"
	"lx/srv/handler/model"

	"github.com/gin-gonic/gin"
)

// ProductHandler 商品处理器
type ProductHandler struct{}

// NewProductHandler 创建商品处理器
func NewProductHandler() *ProductHandler {
	return &ProductHandler{}
}

// CreateProduct 创建商品
func (h *ProductHandler) CreateProduct(c *gin.Context) {
	var req request.CreateProductRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := config.ProductClient.CreateProduct(context.Background(), &proto.CreateProductRequest{
		CategoryId:  int64(req.CategoryID),
		BrandId:     int64(req.BrandID),
		Name:        req.Name,
		Images:      req.Images,
		Description: req.Description,
		Price:       req.Price,
		Stock:       int32(req.Stock),
		Status:      int32(req.Status),
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}

// UpdateProduct 更新商品
func (h *ProductHandler) UpdateProduct(c *gin.Context) {
	id := c.Param("id")
	var req request.UpdateProductRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	req.ID, _ = strconv.ParseInt(id, 10, 64)
	resp, err := config.ProductClient.UpdateProduct(context.Background(), &proto.UpdateProductRequest{
		Id:          req.ID,
		CategoryId:  int64(req.CategoryID),
		BrandId:     int64(req.BrandID),
		Name:        req.Name,
		Images:      req.Images,
		Description: req.Description,
		Price:       req.Price,
		Stock:       int32(req.Stock),
		Status:      int32(req.Status),
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}

// DeleteProduct 删除商品
func (h *ProductHandler) DeleteProduct(c *gin.Context) {
	id := c.Param("id")
	id64, _ := strconv.ParseInt(id, 10, 64)

	resp, err := config.ProductClient.DeleteProduct(context.Background(), &proto.DeleteProductRequest{
		Id: id64,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}

// GetProduct 获取商品详情
func (h *ProductHandler) GetProduct(c *gin.Context) {
	id := c.Param("id")
	id64, _ := strconv.ParseInt(id, 10, 64)

	resp, err := config.ProductClient.GetProduct(context.Background(), &proto.GetProductRequest{
		Id: id64,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}

// ListProduct 商品列表
func (h *ProductHandler) ListProduct(c *gin.Context) {
	req := request.ListProductRequest{}
	c.ShouldBindQuery(&req)

	resp, err := config.ProductClient.ListProduct(context.Background(), &proto.ListProductRequest{
		CategoryId: req.CategoryID,
		BrandId:    req.BrandID,
		Keyword:    req.Keyword,
		Status:     int32(req.Status),
		Page:       int32(req.Page),
		PageSize:   int32(req.PageSize),
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}

// CreateOrder 创建订单
func (h *ProductHandler) CreateOrder(c *gin.Context) {
	var req request.CreateOrderRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 构建订单商品列表
	var list []*proto.OrderItem
	for _, item := range req.List {
		list = append(list, &proto.OrderItem{
			ProductId: item.ProductId,
			Quantity:  item.Quantity,
		})
	}

	resp, err := config.ProductClient.OrderItem(context.Background(), &proto.OrderItemAddRequest{
		UserId:    req.UserId,
		PayType:   req.PayType,
		AddressId: req.AddressId,
		List:      list,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}

// NotifyPay 支付宝异步通知
func (h *ProductHandler) NotifyPay(c *gin.Context) {
	c.Request.ParseForm()
	form := c.Request.PostForm
	fmt.Println("aaa", form)
	tradeStatus := c.Request.PostForm.Get("trade_status")
	outTradeNo := c.Request.PostForm.Get("out_trade_no")

	fmt.Println("tradeStatus:", tradeStatus, "outTradeNo:", outTradeNo)

	if tradeStatus != "TRADE_SUCCESS" {
		c.String(http.StatusOK, "fail")
		return
	}

	if outTradeNo == "" {
		c.String(http.StatusOK, "fail")
		return
	}

	// 1. 查询订单
	var order model.Order
	err := config2.DB.Where("order_sn = ?", outTradeNo).First(&order).Error
	if err != nil {
		c.String(http.StatusOK, "fail")
		return
	}

	// 2. 如果订单已处理，直接返回（幂等性）
	if order.Status == 2 {
		c.String(http.StatusOK, "success")
		return
	}

	// 3. 开启事务
	tx := config2.DB.Begin()

	// 4. 更新订单状态为已支付
	order.Status = 2
	err = tx.Save(&order).Error
	if err != nil {
		tx.Rollback()
		c.String(http.StatusOK, "fail")
		return
	}

	// 5. 查询订单明细，扣减库存
	var orderItems []model.OrderItem
	err = tx.Where("order_id = ?", order.ID).Find(&orderItems).Error
	if err != nil {
		tx.Rollback()
		c.String(http.StatusOK, "fail")
		return
	}

	for _, item := range orderItems {
		var product model.Product
		err = tx.First(&product, item.ProductId).Error
		if err != nil {
			tx.Rollback()
			c.String(http.StatusOK, "fail")
			return
		}
		product.Stock -= int(item.Quantity)
		err = tx.Save(&product).Error
		if err != nil {
			tx.Rollback()
			c.String(http.StatusOK, "fail")
			return
		}
	}

	// 6. 提交事务
	tx.Commit()

	c.String(http.StatusOK, "success")
}
