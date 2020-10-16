package producer

import "os"

var producerQueues = os.Getenv("NEW_NEW_RELIC_PRODUCTION_QUEUES")
var redisConnectionURL = os.Getenv("NEW_NEW_RELC_REDIS_CONNECTION_URL")
