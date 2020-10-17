package main

import (
	"flag"
	"log"

	"github.com/Kareem-Emad/new-new-relic/consumer"
	"github.com/Kareem-Emad/new-new-relic/producer"
	"github.com/Kareem-Emad/new-new-relic/server"
)

const (
	serverMode = "server"
	workerMode = "worker"
)

func runServer() {
	var pm producer.ProductionManager
	pm.InitializeRedisConnection()

	server.Start(pm)
}

func runWorker() {
	var cm consumer.JobBuffer
	cm.InitializeWorker()

	for {
		cm.FetchAndExecute()
	}
}

func main() {
	runModePtr := flag.String("run_mode", serverMode, "specifies whether to run as server/worker")
	flag.Parse()

	switch *runModePtr {
	case serverMode:
		runServer()
		return

	case workerMode:
		runWorker()
		return

	default:
		log.Printf("unrecognized mode %s, expected mode in {%s, %s}", *runModePtr, serverMode, workerMode)
		return
	}
}
