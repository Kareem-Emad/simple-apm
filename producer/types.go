package producer

import "gopkg.in/redis.v5"

// ProductionManager the stucture holding all data associated with redis
type ProductionManager struct {
	redisClient *redis.Client
	queues      []string
}
