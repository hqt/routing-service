package zendesk

import "time"

// Station struct represents the station
type Station struct {
	Code        string
	Name        string
	OpeningDate time.Time
}
