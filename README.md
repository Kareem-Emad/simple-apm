# Simple APM

[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![Build Status:](https://github.com/Kareem-Emad/simple-apm/workflows/Build/badge.svg)](https://github.com/Kareem-Emad/simple-apm/actions)
[![GoReportCard example](https://goreportcard.com/badge/github.com/Kareem-Emad/simple-apm)](https://goreportcard.com/report/Kareem-Emad/simple-apm)

an open source/in-house application performance management server that monitors your services and allows you to perform historical/real-time analysis for your endpoints

## setup

First make sure to init schema in both cassendra and elastic search

```sh
./elastic_init_index.sh
```

in `cqlsh`

```sql
create keyspace simple_apm with replication = {'class': 'SimpleStrategy', 'replication_factor': 1};
use simple_apm;

create table request_info (service_name text, url text, method text, status smallint, response_time int, created_at timestamp, primary key (service_name, created_at, method, url));
```

To run server

```shell
# run one instance of sever to handle http request and push jobs in queues
make run_server
```

To run worker

```shell
# you need two types of workers online at least
make run_worker job_type=ES_SYNC target_queue=sync_es batch_size=5 # elastic search sync worker
make run_worker job_type=DB_WRITE target_queue=cassendra_write batch_size=5 # cassendra DB write worker
```

params:

- `batch_size` number of jobs to handle by worker per one exec
- `target_queue` the job queue worker will listen to
- `job_type` the type of job the spawned worker will handle

## Environment Variables

List of envs needs to be setup before starting the service

- `PRODUCTION_QUEUES` comma separated list of queues to push jobs into from server/producer
- `REDIS_CONNECTION_URL` connection url to setup redis client
- `JWT_SECRET` secret used to sign and verify tokens between sdk/server
- `CASSENDRA_KEY_SPACE` name of the keyspace the tables are stored
- `CASSENDRA_HOSTS` number of cassendra nodes that are under simple-apm
- `SERVER_PORT` port used by http server
- `ES_BULK_INSERTION_CONCURRENCY` max concurrency level use in bulk es index updates

## SDK

- JS: <github.com/Kareem-Emad/simple-apm-express>

## custom SDK

You can build your own custom sdk in whatever lang/way you want, just make sure you

- sign a token by the same secret set here in the envs
- use the same request format

```shell
curl --location --request POST 'http://localhost:5000/requests' \
--header 'Authorization: Bearer jwt_token_placeholder' \
--header 'Content-Type: application/json' \
--data-raw '{
    "url": "https://google.com",
    "http_method": "GET",
    "response_time": 3000,
    "service_name": "service_name",
    "status_code": 200,
    "created_at": "2020-04-02 02:10:01"
}'
```

- make sure you use the same date format as Elastic search is optimzied for this layout specificly as shown below

## Cassandra DB Schema

The database contains one table called `request_info`, containing Fields:

| Field | Data Type| Description |
| --- | --- | --- |
| `service_name` | `text/varchar` | name of the service that contain this endpoint|
| `url` | `text/varchar` | full url of the route in this service |
| `method` | `text/varchar` | type of the request handled by the endpoint |
| `status` | `small_int` | code returned to the requestor |
| `response_time` | `int` | time taken from request recieval in the server to responding to the client |
| `created_at` | `timestamp` | the timestamp when this request was done |

### Indexing

We have two types of indexes:

- `partition index` set to  `service_name` column
- `cluster_index` set to the (`created_at`, `method`, `url`) in the order respectively

Note that accessing the data in cassendra should be optimized to use the index in its order to avoid latencies
So it's better to always start by specifying `service_name`, range of dates `created_at`, http method `method`, and finally the endpoints you want to include in the query result `url`.

## Elastic Search mapping

| Field | Data Type|
| --- | --- |
| `service_name` | `keyword` |
| `url` | `text` |
| `http_method` | `keyword` |
| `status_code` | `short` |
| `response_time` | `long` |
| `created_at` | `date` in format `yyyy-MM-dd HH:mm:ss` |

## Acessing analytics

To access data through elastic search server, there are already some ready to use queries to
fetch important data like `throughput`, `average_response_time`, `min/max_response_time`, `x_percentile`

### throughput

To calculate throughput, we need to use histograms in elastic search.
Following on the assumption that `throughput` is the sucessfull number of requests handled per miute, we can achieve it with a query/aggregate like this one

```shell
curl -X GET "localhost:9200/request_info/_search?pretty" -H 'Content-Type: application/json' -d'
{
    "query": {
      "bool": {
        "must": [ {"match": {"service_name": "YOUR_SERVICE_NAME"}} , {"match":{"url": "SOME_URL"}}, {"match":{"http_method": "GET"}}],
        "filter": {
          "range": {
            "created_at": {
              "gte": "START_DATE IN FORMAT (2020-04-01 02:01:01)",
              "lte": "END_DATE IN FORMAT (2020-04-10 02:01:01)"
            }
          }
        }
      }
    }, "aggs": {
		  "throughput_over_time": {
                "date_histogram": {
                "field": "created_at",
                "fixed_interval": "1m"
            }
          }
    }
}
'
```

Note that you can omit some of the matche queries in the `must` to get calculate the througput curve over your whole service or a specific set of endpoints

Also notice that the interval here is tunable through `fixed_interval` field, meaning you can calculate the throughput of total successfull requests per minute/hour/day/month/etc

### stats(average/min/max/x_percent response time)

To get the full stats for a certain endpoint(s)/service, we can use this query/aggregate:

```shell
curl -X GET "localhost:9200/request_info/_search?pretty" -H 'Content-Type: application/json' -d'
{
    "query": {
      "bool": {
        "must": [ {"match": {"service_name": "YOUR_SERVICE_NAME"}} , {"match":{"url": "SOME_URL"}}, {"match":{"http_method": "POST"}}],
        "filter": {
          "range": {
            "created_at": {
              "gte": "START_DATE",
              "lte": "END_DATE"
            }
          }
        }
      }
    }, "aggs": {
          "request_percentiles": {
              "percentiles": {
                  "field": "response_time" 
              }
          },
          "request_stats":{
              "stats":{
                  "field": "response_time"
              }
          }
    }
}
'
```

Note that you need to specify the period of your search in both curls in the `range` filter to get your stats within a certain timeline (last week/month/...)
