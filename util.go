package zendesk

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"time"
)

// defines all constants
const (
	FullDateLayout = "2 January 2006"
	MonthLayout    = "January 2006"
)

// ParseCSVToStations parse csv file and convert to list of station
func ParseCSVToStations(path string) ([]*Station, error) {
	csvFile, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	r := csv.NewReader(csvFile)

	// remove the header
	_, err = r.Read()
	if err != nil {
		return nil, err
	}

	var stations []*Station
	for {
		record, err := r.Read()
		if err == io.EOF {
			return stations, nil
		}

		if err != nil {
			return nil, err
		}

		if len(record) != 3 {
			return nil, fmt.Errorf("%v must have 3 fields", record)
		}

		date, err := parseDate(record[2])
		if err != nil {
			return nil, err
		}
		station := &Station{
			Code:        record[0],
			Name:        record[1],
			OpeningDate: *date,
		}
		stations = append(stations, station)
	}
}

func parseDate(str string) (*time.Time, error) {
	var t time.Time
	t, err := time.Parse(FullDateLayout, str)
	if err != nil {
		t, err = time.Parse(MonthLayout, str)
		if err != nil {
			return nil, err
		}
	}

	return &t, nil
}
