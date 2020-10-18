package consumer

import (
	"github.com/Kareem-Emad/new-new-relic/dal"
	"github.com/Kareem-Emad/new-new-relic/elasticsearchmanager"
	"gopkg.in/redis.v5"
)

// JobBuffer the stucture holding all data associated with redis
type JobBuffer struct {
	redisClient  *redis.Client
	esClient     *elasticsearchmanager.ESClient
	requestModel *dal.RequestModel
}

const dbWrite = "DB_WRITE"
const esSync = "ES_SYNC"
