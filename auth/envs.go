package auth

import "os"

var jwtSecret = os.Getenv("JWT_SECRET")
