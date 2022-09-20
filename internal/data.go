package alphavantage

import (
	"fmt"
	"sort"
	"time"
)

// full object response from alphavantage api
type Response struct {
	MetaData        MetaData          `json:"Meta Data"`
	DailyTimeSeries map[string]Record `json:"Time Series (Daily)"`
}

// metadata associated with the response
type MetaData struct {
	Information   string `json:"1. Information"`
	Symbol        string `json:"2. Symbol"`
	LastRefreshed string `json:"3. Last Refreshed"`
	OutputSize    string `json:"4. Output Size"`
	TimeZone      string `json:"5. Time Zone"`
}

// daily trading results for a trading symbol
type Record struct {
	Date   time.Time
	Open   float64 `json:"1. open,string"`
	High   float64 `json:"2. high,string"`
	Low    float64 `json:"3. low,string"`
	Close  float64 `json:"4. close,string"`
	Volume int     `json:"5. volume,string"`
}

type TimeSeries []Record

// Timeseries implementations for sort.Interface methods
func (ts TimeSeries) Len() int {
	return len(ts)
}
func (ts TimeSeries) Less(i, j int) bool {
	return ts[i].Date.Before(ts[j].Date)
}
func (ts TimeSeries) Swap(i, j int) {
	ts[i].Date, ts[j].Date = ts[j].Date, ts[i].Date
}

// returns a sorted slice of daily trading results for a trading symbol
// so we can return the records in a legible order and handle missing dates
func (resp *Response) SortTimeSeries() (TimeSeries, error) {
	loc, err := time.LoadLocation(resp.MetaData.TimeZone)
	if err != nil {
		return TimeSeries{}, fmt.Errorf("failed to decode timezone: %w", err)
	}

	// Allocate space for the records to avoid unnecessary reallocation.
	ts := make(TimeSeries, 0, len(resp.DailyTimeSeries))

	for k, v := range resp.DailyTimeSeries {
		var t time.Time
		t, err = time.ParseInLocation("2006-01-02", k, loc)
		if err != nil {
			return TimeSeries{}, err
		}

		v.Date = t

		ts = append(ts, v)
	}

	sort.Sort(sort.Reverse(ts))
	return ts, nil
}
