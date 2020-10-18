package producer

import "os"

var producerQueues = os.Getenv("PRODUCTION_QUEUES")
var redisConnectionURL = os.Getenv("REDIS_CONNECTION_URL")
