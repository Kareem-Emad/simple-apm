package dal

import (
	"fmt"
	"log"
)

const tag = "[DAL]"

// IntializeDBSession starts a new session with cassendra DB
func (rm *RequestModel) IntializeDBSession(dbClient *DBClient) {
	if dbClient != nil {
		rm.dbClient = *dbClient
	} else {
		rm.dbClient = &CassandraDBClient{}
	}

	if err := rm.dbClient.InitSession(); err != nil {
		log.Panicf("%s failed to connect to cassandra cluster | errorLog %s", tag, err)
	}
}

// CreateRequestStats batch based insertion method to create new request state objects in DB
func (rm *RequestModel) CreateRequestStats(requests []RequestStats) bool {
	var queryStatments []string
	for _, req := range requests {
		if ValidateRequestStats(req) == true {
			queryString := "INSERT INTO %s.request_info (service_name, created_at,  method, url, status, response_time)   VALUES ('%s', '%s', '%s', '%s', %d, %d)"
			cqlStatment := fmt.Sprintf(queryString, cassandraKeySpace, req.Service, req.CreatedAt, req.Method, req.URL, req.Status, req.TimeInMilliseconds)

			queryStatments = append(queryStatments, cqlStatment)
		}
	}

	if err := rm.dbClient.BulkInsert(queryStatments); err != nil {
		log.Printf("%s failed to execute batch of request stats insertions | error: %s", tag, err)
		return false
	}
	return true
}
