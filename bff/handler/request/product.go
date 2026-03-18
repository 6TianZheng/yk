package request

// 创建商品请求
type CreateProductRequest struct {
	CategoryID  int64   `json:"category_id"`
	BrandID     int64   `json:"brand_id"`
	Name        string  `json:"name"`
	Images      string  `json:"images"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Stock       int     `json:"stock"`
	Status      int     `json:"status"`
}

// 更新商品请求
type UpdateProductRequest struct {
	ID          int64   `json:"id"`
	CategoryID  int64   `json:"category_id"`
	BrandID     int64   `json:"brand_id"`
	Name        string  `json:"name"`
	Images      string  `json:"images"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Stock       int     `json:"stock"`
	Status      int     `json:"status"`
}

// 删除商品请求
type DeleteProductRequest struct {
	ID int64 `json:"id"`
}

// 获取商品请求
type GetProductRequest struct {
	ID int64 `json:"id"`
}

// 商品列表请求
type ListProductRequest struct {
	CategoryID int64  `json:"category_id"`
	BrandID    int64  `json:"brand_id"`
	Keyword    string `json:"keyword"`
	Status     int    `json:"status"`
	Page       int    `json:"page"`
	PageSize   int    `json:"page_size"`
}

// 订单商品项
type OrderItem struct {
	ProductId int64 `json:"product_id"`
	Quantity  int64 `json:"quantity"`
}

// 创建订单请求
type CreateOrderRequest struct {
	UserId    int64       `json:"user_id"`
	PayType   int64       `json:"pay_type"`
	AddressId int64       `json:"address_id"`
	List      []OrderItem `json:"list"`
}
