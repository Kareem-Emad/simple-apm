package consumer

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/Kareem-Emad/simple-apm/dal"
	"github.com/Kareem-Emad/simple-apm/elasticsearchmanager"
	"gopkg.in/redis.v5"
)

var tag = fmt.Sprintf("[WORKER]")

// InitializeWorker starts a connection with the redis server and inits a job handler
func (jb *JobBuffer) InitializeWorker(targetQueue string, jobType string, batchSize int) {
	jb.targetQueue = targetQueue
	jb.jobType = jobType
	jb.batchSize = batchSize

	tag = fmt.Sprintf("[WORKER|%s]", jb.jobType)

	jb.redisClient = redis.NewClient(&redis.Options{
		Addr: redisConnectionURL,
	})
	// job handler init steps
	switch jobType {

	case dbWrite:
		var rm dal.RequestModel
		rm.IntializeDBSession()

		jb.requestModel = &rm
		return
	case esSync:
		var esClient elasticsearchmanager.ESClient
		esClient.IntializeESClient()

		jb.esClient = &esClient
		return
	}
}

// fetchNewJobs fetches a new batch of jobs from queue
func (jb *JobBuffer) fetchNewJobs() []dal.RequestStats {
	requests := make([]dal.RequestStats, jb.batchSize)

	var currentRequest dal.RequestStats

	log.Printf("%s listening on queue {%s} to fetch job batch of size {%d}", tag, jb.targetQueue, jb.batchSize)

	for idx := range requests {
		res, err := jb.redisClient.BLPop(0, jb.targetQueue).Result()

		if err != nil {
			log.Printf("%s failed to fetch new job from redis queue %s | errorLog: %s", tag, jb.targetQueue, err)
		} else {

			if len(res) == 2 { // [command string, result string]
				err := json.Unmarshal([]byte(res[1]), &currentRequest)

				if err == nil { // yeah we will drop this job because we cannnot parse it
					//now and we won't be able to parse then, it's the same faulty bytes string
					requests[idx] = currentRequest
				} else {
					// let's just log this here for sack of clarity in logs
					log.Fatalf("%s found invalid job bytes string %s, failed to parse", tag, res[1])
				}
			}
		}
	}
	log.Printf("%s collected {%d} jobs from queue", tag, len(requests))
	return requests
}

// executeJobs writes the batch of jobs data fetched from redis into DB
func (jb *JobBuffer) executeJobs(requests []dal.RequestStats) bool {
	switch jb.jobType {

	case dbWrite:
		return jb.requestModel.CreateRequestStats(requests)
	case esSync:
		return jb.esClient.SyncES(requests)

	}

	return false
}

// rollbackReadJobs writes back the jobs into the queue
func (jb *JobBuffer) rollbackReadJobs(requests []dal.RequestStats) {
	for _, req := range requests {

		dataBytes, err := json.Marshal(req)
		if err != nil {
			log.Fatalf("%s Failed to parse back job into bytes | error %s", tag, err)
		} else {

			// push them back on the other side (right side) so they get delayed a bit
			// and not be caught in an infinite loop reading and writing those jobs
			_, err = jb.redisClient.RPush(jb.targetQueue, dataBytes).Result()
			if err != nil {
				log.Fatalf("%s Failed to write back job {%s} to queue | error %s", tag, dataBytes, err)
			}
		}

	}
}

// FetchAndExecute fetches a batch of jobs from redis to execute them
func (jb *JobBuffer) FetchAndExecute() int {
	requests := jb.fetchNewJobs()

	if len(requests) <= 0 { // a lot will say why less but I trust no one xd
		return 0
	}

	status := jb.executeJobs(requests)

	if status == false { // we failed to write those to the database, let's put them back where we found them
		// jb.rollbackReadJobs(requests)
		return -1
	}

	return len(requests)
}
