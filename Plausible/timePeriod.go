package plausible

import "fmt"

// TimeInterval represents an interval of time by which each entry in the query results must be separated by.
type TimeInterval string

const (
	// DateInterval represents a time interval for a particular day.
	// With this time interval, each data point in the result will refer to a particular day.
	DateInterval = TimeInterval("date")
	// MonthInterval represents a time interval of a particular month.
	// With this time interval, each data point in the result will refer to a particular month.
	MonthInterval = TimeInterval("month")
)

func (t *TimeInterval) toQueryArgs() QueryArgs {
	return QueryArgs{
		{Name: "interval", Value: string(*t)},
	}
}

// IsEmpty tells whether the time interval has information.
func (t *TimeInterval) IsEmpty() bool {
	return string(*t) == ""
}
