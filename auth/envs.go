package auth

import "os"

var jwtSecret = os.Getenv("NEW_NEW_RELIC_JWT_SECRET")
