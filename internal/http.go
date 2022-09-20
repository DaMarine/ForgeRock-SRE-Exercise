package alphavantage

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"
)

var url string = fmt.Sprintf("https://www.alphavantage.co/query?apikey=%s&function=TIME_SERIES_DAILY&symbol=%s&outputsize=compact", os.Getenv("APIKEY"), os.Getenv("SYMBOL"))

func Request(ctx context.Context) *http.Response {
	client := http.Client{
		Timeout: 30 * time.Second,
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		panic(err)
	}
	var resp *http.Response
	resp, err = client.Do(req)
	if err != nil {
		panic(err)
	}
	return resp
}
