package zendesk

import (
	"errors"
	"strconv"
	"time"
)

// Station struct represents the station
type Station struct {
	Line        string
	Order       int
	Code        string
	Name        string
	OpeningDate time.Time
}

// NewStation returns a new station
func NewStation(code string, name string, openingDate time.Time) (*Station, error) {
	if len(code) < 3 {
		return nil, errors.New("invalid code format")
	}
	line := code[0:2]
	order, err := strconv.Atoi(code[2:])
	if err != nil {
		return nil, err
	}

	return &Station{
		Line:        line,
		Order:       order,
		Code:        code,
		Name:        name,
		OpeningDate: openingDate,
	}, nil
}

// IsOpen checks if the station is opened at the predefined time or not
func (s *Station) IsOpen(time time.Time) bool {
	return s.OpeningDate.Before(time)
}
