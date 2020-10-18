package dal

import (
	"fmt"
	"log"
	"net/url"
	"strings"

	"github.com/gocql/gocql"
)

var tag = "[DAL]"

// IntializeDBSession starts a new session with cassendra DB
func (rm *RequestModel) IntializeDBSession() {
	cluster := gocql.NewCluster(strings.Split(cassendraHosts, ",")...)

	cluster.Authenticator = gocql.PasswordAuthenticator{
		Username: dbUser,
		Password: dbPassowrd,
	}
	cluster.Keyspace = cassandraKeySpace
	cluster.Consistency = gocql.Quorum

	session, err := cluster.CreateSession()

	if err != nil {
		log.Panicf("%s failed to connect to cassandra cluster | errorLog %s", tag, err)
	}

	rm.dbSession = session
}

// CreateRequestStats batch based insertion method to create new request state objects in DB
func (rm *RequestModel) CreateRequestStats(requests []RequestStats) bool {
	batch := rm.dbSession.NewBatch(gocql.LoggedBatch)

	for _, req := range requests {
		if ValidateRequestStats(req) == true {
			queryString := "INSERT INTO %s.request_info (service_name, created_at,  method, url, status, response_time)   VALUES ('%s', '%s', '%s', '%s', %d, %d)"

			cqlStatment := fmt.Sprintf(queryString, cassandraKeySpace, req.Service, req.CreatedAt, req.Method, req.URL, req.Status, req.TimeInMilliseconds)
			log.Printf("%s adding query to batch: %s", tag, cqlStatment)

			batch.Query(cqlStatment)
		}
	}

	if err := rm.dbSession.ExecuteBatch(batch); err != nil {
		log.Printf("%s failed to execute batch of request stats insertions | error: %s", tag, err)
		return false
	}
	return true
}

// ValidateRequestStats checks that the data in this request object is valid to be inserted in DB
func ValidateRequestStats(req RequestStats) bool {
	validStatus := req.Status >= 200 && req.Status <= 500
	validResponseTime := req.TimeInMilliseconds > 0

	url, err := url.ParseRequestURI(req.URL)
	validURL := (err == nil && url.Scheme != "" && url.Host != "")

	validHTTPMethods := []string{"GET", "POST", "PATCH", "HEAD", "DELETE"}
	validMethod := isInArray(validHTTPMethods, strings.ToUpper(req.Method))

	validServiceName := (req.Service != "")

	// time.Now().Format(time.RFC3339)
	return validStatus && validResponseTime && validURL && validMethod && validServiceName && req.CreatedAt != ""
}
