package dal

import (
	"net/url"
	"strings"
)

var validHTTPMethods = []string{"GET", "POST", "PATCH", "HEAD", "DELETE"}

// ValidateRequestStats checks that the data in this request object is valid to be inserted in DB
func ValidateRequestStats(req RequestStats) bool {
	validStatus := req.Status >= 200 && req.Status <= 500
	validResponseTime := req.TimeInMilliseconds > 0

	url, err := url.ParseRequestURI(req.URL)
	validURL := (err == nil && url.Scheme != "" && url.Host != "")

	validMethod := isInArray(validHTTPMethods, strings.ToUpper(req.Method))

	validServiceName := (req.Service != "")

	// time.Now().Format(time.RFC3339)
	return validStatus && validResponseTime && validURL && validMethod && validServiceName && req.CreatedAt != ""
}
