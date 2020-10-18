package elasticsearchmanager

import (
	"github.com/Kareem-Emad/simple-apm/envreader"
)

var esBulkConcurrency = envreader.ReadEnvAsInt("ES_BULK_INSERTION_CONCURRENCY", 2)
