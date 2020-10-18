package consumer

import (
	"os"
)

var redisConnectionURL = os.Getenv("REDIS_CONNECTION_URL")
