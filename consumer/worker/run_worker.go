package main

import (
	"github.com/Kareem-Emad/new-new-relic/consumer"
)

func main() {
	var cm consumer.JobBuffer
	cm.InitializeWorker()

	for {
		cm.FetchAndExecute()
	}
}
