package main

import (
	"github.com/Kareem-Emad/new-new-relic/producer"
	"github.com/Kareem-Emad/new-new-relic/server"
)

func main() {
	var pm producer.ProductionManager
	pm.InitializeRedisConnection()

	server.Start(pm)
}
