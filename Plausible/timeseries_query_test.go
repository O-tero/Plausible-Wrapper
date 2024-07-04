package plausible

import (
	"os"
	"strings"
	"testing"
)

func TestUnitValidateTimeseriesQuery(t *testing.T) {
	tests := []struct {
		name    string
		query   TimeseriesQuery
		isValid bool
	}{
		{
			name: "valid timeseries query",
			query: TimeseriesQuery{
				Period:   DayPeriod(),
				Filters:  NewFilter().ByVisitOs("Windows"),
				Metrics:  AllMetrics(),
				Interval: DateInterval,
			},
			isValid: true,
		},
		{
			name: "invalid timeseries due to missing time period",
			query: TimeseriesQuery{
				Filters:  NewFilter().ByVisitOs("Windows"),
				Metrics:  AllMetrics(),
				Interval: DateInterval,
			},
			isValid: false,
		},
	}

	for _, test := range tests {
		valid, _ := test.query.Validate()
		if valid && !test.isValid {
			t.Fatalf("test '%s' is valid, but was expected to fail", test.name)
		}
		if !valid && test.isValid {
			t.Fatalf("test '%s' is invalid, but was expected to succeed", test.name)
		}
	}

}

func TestUnitToQueryArgsTimeseriesQuery(t *testing.T) {
	tests := []struct {
		name              string
		query             TimeseriesQuery
		expectedQueryArgs QueryArgs
		isValid           bool
	}{
		{
			name: "valid breakdown query",
			query: TimeseriesQuery{
				Period:   DayPeriod(),
				Filters:  NewFilter().ByVisitOs("Windows"),
				Metrics:  AllMetrics(),
				Interval: DateInterval,
			},
			expectedQueryArgs: QueryArgs{
				QueryArg{Name: "period", Value: "day"},
				QueryArg{Name: "filters", Value: "visit:os==Windows"},
				QueryArg{Name: "metrics", Value: "visitors,pageviews,bounce_rate,visit_duration,visits"},
				QueryArg{Name: "interval", Value: "date"},
			},
			isValid: true,
		},