package stock

import (
	"context"
	"encoding/json"
	"os"
	alphavantage "stockticker/internal"
	"strconv"
)

type Server struct {
	UnimplementedTickerServer
}

func (s *Server) GetAvgPrice(ctx context.Context, in *StockRequest) (*StockResponse, error) {

	periodOfTime, err := strconv.Atoi(os.Getenv("NDAYS"))
	if err != nil {
		return nil, err
	}

	// request data from alphavantage and decode the response
	resp := alphavantage.Request(ctx)
	defer resp.Body.Close()
	decoder := json.NewDecoder(resp.Body)
	decoder.DisallowUnknownFields()
	var data alphavantage.Response
	err = decoder.Decode(&data)
	if err != nil {
		panic(err)
	}

	// sort the data by date, before we calculate the average and number of days of data to return
	var series alphavantage.TimeSeries
	series, err = data.SortTimeSeries()
	if err != nil {
		panic(err)
	}

	var dailyDataArr []*DailyData
	var avgClosePrice float64

	totalClosePrices := 0.0
	totalTradingDays := 0.0
	remainingDays := int(periodOfTime)
	previousDay := series[0].Date.AddDate(0, 0, 1)
	for _, record := range series[:periodOfTime] {

		// account for any weekends and holidays where there is no data
		// by calculating the number of days between the previous day and the current day
		timeDifference := record.Date.Sub(previousDay)
		remainingDays += int(timeDifference.Hours() / 24)
		if remainingDays >= 0 {

			// add data to the array of daily data and avg closing price calculations
			data := &DailyData{
				Date:   record.Date.String(),
				Open:   record.Open,
				High:   record.High,
				Low:    record.Low,
				Close:  record.Close,
				Volume: int32(record.Volume),
			}
			dailyDataArr = append(dailyDataArr, data)
			totalClosePrices += record.Close
			previousDay = record.Date
			totalTradingDays++
		} else {
			break
		}
	}
	avgClosePrice = (totalClosePrices / totalTradingDays)
	return &StockResponse{AvgClosePrice: avgClosePrice, DailyData: dailyDataArr}, nil
}
