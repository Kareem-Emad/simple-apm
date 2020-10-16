package server

// Message structure expected from publishers
type Message struct {
	URL                string
	Method             string
	TimeInMilliseconds int
}
