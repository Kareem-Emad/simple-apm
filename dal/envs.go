package dal

import "os"

var dbUser = os.Getenv("DB_USER")
var dbPassowrd = os.Getenv("DB_PASSWORD")
var cassandraKeySpace = os.Getenv("CASSENDRA_KEY_SPACE")
var cassendraHosts = os.Getenv("CASSENDRA_HOSTS")
