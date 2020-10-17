package dal

import "github.com/gocql/gocql"

// RequestModel structure holding an orm like wrapper around RequestStats table
type RequestModel struct {
	dbSession *gocql.Session
}

// RequestStats structure holding data for the request stats table
type RequestStats struct {
	URL                string `json:"url"`
	Method             string `json:"http_method"`
	TimeInMilliseconds int    `json:"response_time"`
	Service            string `json:"service_name"`
	Status             uint16 `json:"status_code"`
}
