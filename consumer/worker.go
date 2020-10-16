package consumer

import (
	"log"
	"strconv"

	"github.com/Kareem-Emad/new-new-relic/dal"
	elasticSearch "github.com/Kareem-Emad/new-new-relic/elastic_search"
	"gopkg.in/redis.v5"
)

var tag = "[WORKER]"

// InitializeWorker starts a connection with the redis server and inits a job handler
func (jb *JobBuffer) InitializeWorker() {
	jb.redisClient = redis.NewClient(&redis.Options{
		Addr: redisConnectionURL,
	})
	// job handler init steps
	switch jobType {

	case dbWrite:
		var rm dal.RequestModel
		rm.IntializeDBSession()

		jb.requestModel = &rm
	case esSync:
	}
}

// fetchNewJobs fetches a new batch of jobs from queue
func (jb *JobBuffer) fetchNewJobs() []string {
	bs, _ := strconv.Atoi(batchSize)
	requests := make([]string, bs)

	for idx := range requests {
		res, err := jb.redisClient.BLPop(0, queueName).Result()

		if err != nil {
			log.Printf("%s failed to fetch new job from redis queue %s | errorLog: %s", tag, queueName, err)
		} else {
			if len(res) == 2 { // command string + result string in the total result array
				requests[idx] = res[1]
			}
		}
	}

	return requests
}

// executeJobs writes the batch of jobs data fetched from redis into DB
func (jb *JobBuffer) executeJobs(requests []string) bool {
	switch jobType {

	case dbWrite:
		return jb.requestModel.CreateRequestStats(requests)
	case esSync:
		return elasticSearch.SyncES(requests)

	}

	return false
}

// rollbackReadJobs writes back the jobs into the queue
func (jb *JobBuffer) rollbackReadJobs(requests []string) {
	// push them back on the other side (right side) so they get delayed a bit
	// and not be caught in an infinite loop reading and writing those jobs
	_, err := jb.redisClient.RPush(queueName, requests).Result()

	if err != nil {
		log.Fatalf("%s Failed to write back jobs to queue | error %s", tag, err)
	}
}

// FetchAndExecute fetches a batch of jobs from redis to execute them
func (jb *JobBuffer) FetchAndExecute() int {
	requests := jb.fetchNewJobs()

	if len(requests) <= 0 {
		return 0
	}

	status := jb.executeJobs(requests)

	if status == false { // we failed to write those to the database, let's put them back where we found them
		jb.rollbackReadJobs(requests)
		return -1
	}

	return len(requests)
}
