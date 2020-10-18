package main

import (
	"flag"
	"log"

	"github.com/Kareem-Emad/simple-apm/consumer"
	"github.com/Kareem-Emad/simple-apm/producer"
	"github.com/Kareem-Emad/simple-apm/server"
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

func runWorker(targetQueue string, jobType string, batchSize int) {
	var cm consumer.JobBuffer
	cm.InitializeWorker(targetQueue, jobType, batchSize)

	for {
		cm.FetchAndExecute()
	}
}

func main() {
	runModePtr := flag.String("run_mode", serverMode, "specifies whether to run as server/worker")

	targetQueuePtr := flag.String("target_queue", "none", "queue for worker to pull data from")
	jobTypePtr := flag.String("job_type", "none", "type of job the worker will handle")
	batchSizePtr := flag.Int("batch_size", 1, "Number of jobs to handle at one exec")

	flag.Parse()

	switch *runModePtr {
	case serverMode:
		runServer()
		return

	case workerMode:
		runWorker(*targetQueuePtr, *jobTypePtr, *batchSizePtr)
		return

	default:
		log.Printf("unrecognized mode %s, expected mode in {%s, %s}", *runModePtr, serverMode, workerMode)
		return
	}
}
