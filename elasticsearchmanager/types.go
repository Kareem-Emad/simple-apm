package elasticsearchmanager

import elasticsearch "github.com/elastic/go-elasticsearch/v7"

// ESClient structure holding es client instance
type ESClient struct {
	es *elasticsearch.Client
}
