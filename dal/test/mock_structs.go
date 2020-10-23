package daltest

import (
	"testing"
)

// MockDBClient mock
type MockDBClient struct {
	TestObject        *testing.T
	isInit            bool
	ExpectedBulkQuery []string
}

// InitSession mock
func (mdbc *MockDBClient) InitSession() error {
	mdbc.isInit = true
	return nil
}

// BulkInsert mock
func (mdbc *MockDBClient) BulkInsert(insertQueries []string) error {
	if mdbc.isInit != true {
		mdbc.TestObject.Fatal("DB Client should be initialized before use")
	}

	for idx := range mdbc.ExpectedBulkQuery {
		if mdbc.ExpectedBulkQuery[idx] != insertQueries[idx] {
			mdbc.TestObject.Fatalf("expected value %s at position %d | found %s in array of bulk insert queries",
				mdbc.ExpectedBulkQuery[idx], idx, insertQueries[idx])
		}
	}

	return nil
}
