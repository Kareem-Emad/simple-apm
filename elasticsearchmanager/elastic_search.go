package elasticsearchmanager

import (
	"bytes"
	"context"
	"encoding/json"
	"log"

	"github.com/Kareem-Emad/simple-apm/dal"
	elasticsearch "github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esutil"
)

const tag = "[ElasticSearchClient]"
const esIndexName = "request_info"

// IntializeESClient inits a new es client instance
func (esc *ESClient) IntializeESClient() {
	es, err := elasticsearch.NewDefaultClient()

	if err != nil {
		log.Panicf("%s Failed to initalize es client | error %s", tag, err)
	}

	esc.es = es
}

// SyncES syncs a batch of new request stats records to elastic search index
func (esc *ESClient) SyncES(requests []dal.RequestStats) bool {
	bi, err := esutil.NewBulkIndexer(esutil.BulkIndexerConfig{
		Index:      esIndexName,       // The default index name
		Client:     esc.es,            // The Elasticsearch client
		NumWorkers: esBulkConcurrency, // The number of worker goroutines
	})

	if err != nil {
		log.Fatalf("%s failed to intialize a bulk insertion proc in ElasticSearch | error %s", tag, err)
		return false
	}

	for _, req := range requests {

		data, err := json.Marshal(req)
		if err != nil {
			log.Fatalf("%s Failed to encode request for bulk es insert | error %s", tag, err)
		}

		// Add an item to the BulkIndexer
		err = bi.Add(
			context.Background(),
			esutil.BulkIndexerItem{
				// Action field configures the operation to perform (index, create, delete, update)
				Action: "index",
				// Body is an `io.Reader` with the payload
				Body: bytes.NewReader(data),
				// OnFailure is called for each failed operation
				OnFailure: func(ctx context.Context, item esutil.BulkIndexerItem, res esutil.BulkIndexerResponseItem, err error) {
					if err != nil {
						log.Fatalf("%s Internal failure in bulk indexer | error %s", tag, err)
					} else {
						log.Fatalf("%s Internal failure in bulk indexer | error type{%s}=> %s", tag, res.Error.Type, res.Error.Reason)
					}
				},
			},
		)
	}

	if err := bi.Close(context.Background()); err != nil {
		log.Fatalf("%s failed to close bulk es inserter session | error %s", tag, err)
		return false
	}
	return true
}
