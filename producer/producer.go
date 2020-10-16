package producer

import (
	"fmt"
	"log"
	"strings"

	"gopkg.in/redis.v5"
)

var tag = "[NEW_NEW_RELIC_Producer]"

// InitializeRedisConnection starts a connection with the redis server
func (pm *ProductionManager) InitializeRedisConnection() {
	pm.redisClient = redis.NewClient(&redis.Options{
		Addr: redisConnectionURL,
	})

	pm.queues = strings.Split(producerQueues, ",")
}

// PushJobInQueues inserts the new http request stats in redis queues to be stored to db
func (pm *ProductionManager) PushJobInQueues(method string, url string, timeInMilliseconds int) bool {
	stringfiedMessage := fmt.Sprintf("%s|%s|%d", method, url, timeInMilliseconds)

	for _, qkey := range pm.queues {
		err := pm.redisClient.LPush(qkey, stringfiedMessage).Err()

		if err != nil {
			log.Fatalf("%s Failed to insert job [%s] into queue %s | errorlog: %s", tag, stringfiedMessage, qkey, err)
			return false
		}
	}

	return true
}
