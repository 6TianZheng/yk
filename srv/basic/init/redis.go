package initpkg

import (
	"fmt"
	"lx/srv/basic/config"

	"github.com/go-redis/redis/v8"
)

func InitCache() {
	r := config.GlobalConfig.Redis

	Addr := fmt.Sprintf("%s:%d", r.Host, r.Port)
	fmt.Println(Addr)
	config.Rdb = redis.NewClient(&redis.Options{
		Addr:     Addr,
		Password: r.Password, // no password set
		DB:       0,          // use default DB
	})
	err := config.Rdb.Ping(config.Ctx).Err()
	if err != nil {
		return
	}
	fmt.Println("redis连接成功")

}
