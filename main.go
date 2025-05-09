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

func main() {
	_ = godotenv.Load()
	config := LoadConfig()

	cryptos, err := FetchTopCryptos(10, FetchRSI, config)
	if err != nil {
		log.Fatalf("Error fetching top cryptos: %v", err)
	}

	for i, crypto := range cryptos {
		rsi, err := FetchRSI(crypto.Symbol, config.Interval, config.RSIPeriod)
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
	table.SetHeader([]string{"#", "Symbol", fmt.Sprintf("RSI(%d)", config.RSIPeriod), "Market Cap (USD)", "Market Cap Rank"})
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
			strconv.Itoa(crypto.Rank),
		})
	}
	table.Render()
}
