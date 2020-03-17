package routing

import (
	"time"
)

// DayType defines day type
type DayType int

// all constants
const (
	PeakTime DayType = iota
	NonPeakTime
	NightTime
)

// TimeConstraint defines time constraint for each station
type TimeConstraint struct {
	PeakTime    int
	NonPeakTime int
	NightTime   int
}

var timeStations = map[string]TimeConstraint{
	"NS":          {12, 10, 10},
	"NE":          {12, 10, 10},
	"DT":          {10, 8, 0},
	"TE":          {10, 8, 8},
	"CG":          {10, 10, 0},
	"CE":          {10, 10, 0},
	"EW":          {10, 10, 10},
	"CC":          {12, 10, 10},
	"CHANGE_LINE": {15, 10, 10},
}

func getDayTypeOfTime(timeObject time.Time) DayType {
	hour := timeObject.Hour()
	weekday := timeObject.Weekday()

	if hour >= 22 || hour <= 6 {
		return NightTime
	}

	if weekday >= time.Monday && weekday <= time.Friday {
		if (hour >= 6 && hour <= 9) || (hour >= 18 && hour <= 21) {
			return PeakTime
		}
	}

	return NonPeakTime
}

func getTimeBetweenStation(first *Station, second *Station, dayType DayType) int {
	var timeConfig TimeConstraint
	if first.Line == second.Line {
		timeConfig = timeStations[first.Line]
	} else {
		timeConfig = timeStations["CHANGE_LINE"]
	}

	if dayType == PeakTime {
		return timeConfig.PeakTime
	} else if dayType == NightTime {
		return timeConfig.NightTime
	} else {
		return timeConfig.NonPeakTime
	}
}
