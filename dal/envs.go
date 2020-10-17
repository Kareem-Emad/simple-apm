package dal

import "os"

var dbUser = os.Getenv("NEW_NEW_RELIC_DB_USER")
var dbPassowrd = os.Getenv("NEW_NEW_RELC_DB_PASSWORD")
var cassandraKeySpace = os.Getenv("NEW_NEW_RELIC_CASSENDRA_KEY_SPACE")
var cassendraHosts = os.Getenv("NEW_NEW_RELIC_CASSENDRA_HOSTS")
