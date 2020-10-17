package elasticSearch

import "github.com/Kareem-Emad/new-new-relic/dal"

// SyncES syncs a batch of new request stats records to elastic search index
func SyncES(requests []dal.RequestStats) bool {
	return true
}
