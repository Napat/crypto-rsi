package main

import (
	"os"
	"strconv"
)

type Config struct {
	Interval  string
	RSIPeriod int
}

func LoadConfig() Config {
	interval := os.Getenv("INTERVAL")
	if interval == "" {
		interval = "15m"
	}
	rsiPeriodStr := os.Getenv("RSI_PERIOD")
	if rsiPeriodStr == "" {
		rsiPeriodStr = "14"
	}
	rsiPeriod, err := strconv.Atoi(rsiPeriodStr)
	if err != nil {
		rsiPeriod = 14
	}
	return Config{
		Interval:  interval,
		RSIPeriod: rsiPeriod,
	}
}
