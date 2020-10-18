package consumer

import (
	"os"

	"github.com/Kareem-Emad/new-new-relic/envreader"
)

var batchSize = envreader.ReadEnvAsInt("WRITE_BATCH_SIZE", 1)
var redisConnectionURL = os.Getenv("REDIS_CONNECTION_URL")
var queueName = os.Getenv("QUEUE_NAME")
var jobType = os.Getenv("JOB_TYPE")
