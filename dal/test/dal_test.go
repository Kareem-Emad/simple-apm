package daltest

import (
	"testing"

	"github.com/Kareem-Emad/simple-apm/dal"
)

// TestBulkInsertionQueryBuild should extract info from requests array
// and construct proper insert queries as array
func TestBulkInsertionQueryBuild(t *testing.T) {
	var mockDB dal.DBClient

	testMockDB := MockDBClient{}
	// let's do some plumbing here before we send away our mock
	testMockDB.ExpectedBulkQuery = insertStatments
	testMockDB.TestObject = t

	mockDB = &testMockDB

	var rm dal.RequestModel
	rm.IntializeDBSession(&mockDB)

	rm.CreateRequestStats(requests)
}
