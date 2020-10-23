package dal

import (
	"strings"

	"github.com/gocql/gocql"
)

// InitSession starts a new db session and returns a flag for
// the status of the operation
func (cdbc *CassandraDBClient) InitSession() error {
	cluster := gocql.NewCluster(strings.Split(cassendraHosts, ",")...)

	cluster.Authenticator = gocql.PasswordAuthenticator{
		Username: dbUser,
		Password: dbPassowrd,
	}
	cluster.Keyspace = cassandraKeySpace
	cluster.Consistency = gocql.Quorum

	session, err := cluster.CreateSession()

	cdbc.dbSession = session
	return err
}

// BulkInsert does a bulk/batch insert of records into the DB
func (cdbc *CassandraDBClient) BulkInsert(insertQueries []string) error {
	batch := cdbc.dbSession.NewBatch(gocql.LoggedBatch)

	for _, query := range insertQueries {
		batch.Query(query)
	}
	err := cdbc.dbSession.ExecuteBatch(batch)

	return err
}
