package dal

import (
	"fmt"
	"log"
	"strings"

	"github.com/gocql/gocql"
)

var tag = "[DAL]"

// IntializeDBSession starts a new session with cassendra DB
func (rm *RequestModel) IntializeDBSession() {
	cluster := gocql.NewCluster("192.168.1.1", "192.168.1.2", "192.168.1.3")
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
func (rm *RequestModel) CreateRequestStats(requests []string) bool {
	cqlStatment := "BEGIN BATCH\n"

	for _, req := range requests {
		reqFields := strings.Split(req, ",")

		queryString := "INSERT INTO %s.requests (url, method, status, response_time) VALUES (%s, %s, %s, %s) IF NOT EXISTS;"
		cqlStatment += fmt.Sprintf(queryString, dbName, reqFields[0], reqFields[1], reqFields[2], reqFields[3])
	}

	cqlStatment += "\nAPPLY BATCH;"

	err := rm.dbSession.Query(cqlStatment)
	if err != nil {
		log.Printf("%s failed to execute batch of insertions | error: %s", tag, err)
		return false
	}
	return true
}
