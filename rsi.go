package main

import (
	"context"
	"fmt"
	"strconv"
	"time"

	binance "github.com/adshao/go-binance/v2"
	talib "github.com/markcheno/go-talib"
)

// FetchRSI fetches RSI for a given symbol, interval, and period from Binance.
func FetchRSI(symbol, interval string, period int) (float64, error) {
	client := binance.NewClient("", "")

	endTime := time.Now().UnixMilli()
	var duration int64

	switch interval {
	case "1m":
		duration = 60 * 1000
	case "5m":
		duration = 5 * 60 * 1000
	case "15m":
		duration = 15 * 60 * 1000
	case "1h":
		duration = 60 * 60 * 1000
	case "4h":
		duration = 4 * 60 * 60 * 1000
	case "1d":
		duration = 24 * 60 * 60 * 1000
	case "1w":
		duration = 7 * 24 * 60 * 60 * 1000
	case "1M":
		duration = 30 * 24 * 60 * 60 * 1000
	default:
		return 0, fmt.Errorf("invalid interval: %s", interval)
	}

	startTime := endTime - int64(200*duration)

	klines, err := client.NewKlinesService().
		Symbol(symbol).
		Interval(interval).
		StartTime(startTime).
		EndTime(endTime).
		Do(context.Background())
	if err != nil {
		return 0, err
	}

	if len(klines) < period {
		return 0, fmt.Errorf("not enough data for %s", symbol)
	}

	closePrices := make([]float64, len(klines))
	for i, k := range klines {
		closePrice, _ := strconv.ParseFloat(k.Close, 64)
		closePrices[i] = closePrice
	}

	rsi := talib.Rsi(closePrices, period)
	return rsi[len(rsi)-1], nil
}
