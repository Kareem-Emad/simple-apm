package producer

import (
	"encoding/json"
	"log"
	"strings"

	"github.com/Kareem-Emad/new-new-relic/dal"
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

// PushJobInQueues inserts the new http request stats in all redis queues to be stored to db
func (pm *ProductionManager) PushJobInQueues(requestStats dal.RequestStats) bool {
	dataBytes, err := json.Marshal(requestStats)

	if err != nil {
		log.Fatalf("%s Failed to seralize job into bytes | error: %s", tag, err)
		return false
	}

	for _, qkey := range pm.queues {
		err := pm.redisClient.LPush(qkey, dataBytes).Err()

		if err != nil {
			log.Fatalf("%s Failed to insert job into queue %s | error: %s", tag, qkey, err)
			return false
		}

	}

	return true
}
