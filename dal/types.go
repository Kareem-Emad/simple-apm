package dal

import "github.com/gocql/gocql"

// DBClient Contract defining the required functionalities
// that need to be avaiable in any DB client to work with other components
// using the DB
type DBClient interface {
	// InitSession starts a new db session and returns a flag for
	// the status of the operation
	InitSession() error
	// BulkInsert does a bulk/batch insert of records into the DB
	BulkInsert([]string) error
}

// RequestModel structure holding an orm like wrapper around RequestStats table
type RequestModel struct {
	dbClient DBClient
}

// RequestStats structure holding data for the request stats table
type RequestStats struct {
	URL                string `json:"url"`
	Method             string `json:"http_method"`
	TimeInMilliseconds int    `json:"response_time"`
	Service            string `json:"service_name"`
	Status             uint16 `json:"status_code"`
	CreatedAt          string `json:"created_at"`
}

// CassandraDBClient an implementation for DBClient wrapping
// cassandra DB client
type CassandraDBClient struct {
	dbSession *gocql.Session
}
