package consumer

import (
	"github.com/Kareem-Emad/new-new-relic/dal"
	"gopkg.in/redis.v5"
)

// JobBuffer the stucture holding all data associated with redis
type JobBuffer struct {
	redisClient  *redis.Client
	requestModel *dal.RequestModel
}

const dbWrite = "DB_WRITE"
const esSync = "ES_SYNC"
