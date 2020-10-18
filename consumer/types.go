package consumer

import (
	"github.com/Kareem-Emad/simple-apm/dal"
	"github.com/Kareem-Emad/simple-apm/elasticsearchmanager"
	"gopkg.in/redis.v5"
)

// JobBuffer the stucture holding all data associated with redis
type JobBuffer struct {
	redisClient  *redis.Client
	esClient     *elasticsearchmanager.ESClient
	requestModel *dal.RequestModel
	targetQueue  string
	jobType      string
	batchSize    int
}

const dbWrite = "DB_WRITE"
const esSync = "ES_SYNC"
