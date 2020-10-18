package consumer

import (
	"os"

	"github.com/Kareem-Emad/new-new-relic/envreader"
)

var batchSize = envreader.ReadEnvAsInt("NEW_NEW_RELIC_WRITE_BATCH_SIZE", 1)
var redisConnectionURL = os.Getenv("NEW_NEW_RELC_REDIS_CONNECTION_URL")
var queueName = os.Getenv("NEW_NEW_RELIC_QUEUE_NAME")
var jobType = os.Getenv("NEW_NEW_RELIC_JOB_TYPE")
