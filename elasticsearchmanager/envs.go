package elasticsearchmanager

import (
	"github.com/Kareem-Emad/new-new-relic/envreader"
)

var esBulkConcurrency = envreader.ReadEnvAsInt("ES_BULK_INSERTION_CONCURRENCY", 2)
