package router

import (
	"lx/bff/handler/service"

	"github.com/gin-gonic/gin"
)

var productHandler = service.NewProductHandler()

func Router() *gin.Engine {
	r := gin.Default()

	// 商品路由
	product := r.Group("/product")
	{
		product.POST("", productHandler.CreateProduct)
		product.PUT("/:id", productHandler.UpdateProduct)
		product.DELETE("/:id", productHandler.DeleteProduct)
		product.GET("/:id", productHandler.GetProduct)
		product.GET("/list", productHandler.ListProduct)
	}

	// 订单路由
	order := r.Group("/order")
	{
		order.POST("", productHandler.CreateOrder)
	}

	// 支付宝异步通知（兼容 GET 和 POST）
	r.POST("/notify/pay", productHandler.NotifyPay)

	return r
}
