package initpkg

import (
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"lx/srv/basic/config"
	"lx/srv/handler/model"
)

func init() {
	ViperInit()
	NacosInit()
	MysqlInit()
	InitCache()
	ConsulInit()
}

func ViperInit() {
	viper.SetConfigFile("../../../config.yaml")

	err := viper.ReadInConfig()
	if err != nil {
		panic("配置读取失败: " + err.Error())
	}

	err = viper.Unmarshal(&config.GlobalConfig)
	if err != nil {
		panic("配置解析失败: " + err.Error())
	}
	fmt.Println(config.GlobalConfig)
}

func MysqlInit() {
	var err error
	data := config.GlobalConfig.Mysql
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		data.User,
		data.Password,
		data.Host,
		data.Port,
		data.Database)
	config.DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("数据库连接失败: " + err.Error())
	}
	fmt.Println("数据库连接成功")

	err = config.DB.AutoMigrate(&model.Inventory{}, &model.Address{}, &model.DeliveryStaff{}, &model.InventoryLog{}, &model.Logistics{},
		model.LogisticsTrack{}, &model.MemberAddress{}, &model.MemberCart{}, &model.MemberLevel{}, &model.Order{}, &model.OrderItem{},
		&model.ProductCategory{}, &model.Product{})
	if err != nil {
		panic("数据库迁移失败: " + err.Error())
	}
	fmt.Println("数据库迁移成功")
}
