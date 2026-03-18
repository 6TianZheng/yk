package main

import (
	"lx/srv/basic/RabbitMQ"
	_ "lx/srv/basic/init"
)

func main() {
	RabbitMQ.SendStockDeductMsg("2", 5)
}
