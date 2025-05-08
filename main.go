package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/olekukonko/tablewriter"
)

func loadEnv() {
	_ = godotenv.Load()
	if os.Getenv("INTERVAL") == "" {
		os.Setenv("INTERVAL", "15m")
	}
	if os.Getenv("RSI_PERIOD") == "" {
		os.Setenv("RSI_PERIOD", "14")
	}
}

func main() {
	loadEnv()

	cryptos, err := FetchTopCryptos(10, FetchRSI)
	if err != nil {
		log.Fatalf("Error fetching top cryptos: %v", err)
	}
	interval := os.Getenv("INTERVAL")
	rsiPeriod, err := strconv.Atoi(os.Getenv("RSI_PERIOD"))
	if err != nil {
		log.Fatalf("Invalid RSI_PERIOD: %v", err)
	}

	for i, crypto := range cryptos {
		rsi, err := FetchRSI(crypto.Symbol, interval, rsiPeriod)
		if err != nil {
			log.Printf("Error fetching RSI for %s: %v", crypto.Symbol, err)
			continue
		}
		cryptos[i].RSI = rsi
	}

	sort.Slice(cryptos, func(i, j int) bool {
		return cryptos[i].RSI > cryptos[j].RSI
	})

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"#", "Symbol", "RSI", "Market Cap (USD)", "Market Cap Rank"})
	table.SetBorder(true)
	table.SetColumnAlignment([]int{tablewriter.ALIGN_RIGHT, tablewriter.ALIGN_CENTER, tablewriter.ALIGN_RIGHT, tablewriter.ALIGN_RIGHT, tablewriter.ALIGN_RIGHT})

	for i, crypto := range cryptos {
		formattedMarketCap := strconv.FormatFloat(crypto.MarketCap, 'f', 0, 64)
		formattedMarketCapWithComma := FormatWithComma(formattedMarketCap)
		table.Append([]string{
			fmt.Sprintf("%d", i+1),
			crypto.Symbol,
			fmt.Sprintf("%.2f", crypto.RSI),
			formattedMarketCapWithComma,
			strconv.Itoa(crypto.Rank), // Show market cap rank
		})
	}
	table.Render()
}
