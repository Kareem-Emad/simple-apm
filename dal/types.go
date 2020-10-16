package dal

import "github.com/gocql/gocql"

type RequestModel struct {
	dbSession *gocql.Session
}
