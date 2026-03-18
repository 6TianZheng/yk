package main

import (
	"lx/srv/basic/RabbitMQ"
	_ "lx/srv/basic/init"
)

func main() {
	RabbitMQ.ConsumeStockDeduct()
}
