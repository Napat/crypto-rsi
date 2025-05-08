package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

// fetchTopCryptos fetches the top cryptocurrencies by market cap, skipping stablecoins and unsupported symbols.
func FetchTopCryptos(limit int, fetchRSIFunc func(symbol, interval string, period int) (float64, error)) ([]CryptoRSI, error) {
	url := fmt.Sprintf("https://api.coingecko.com/api/v3/coins/markets?vs_currency=usd&order=market_cap_desc&per_page=%d&page=1&sparkline=false", limit*2)
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch top cryptos: %v", err)
	}
	defer resp.Body.Close()

	var data []struct {
		Symbol    string  `json:"symbol"`
		MarketCap float64 `json:"market_cap"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, fmt.Errorf("failed to decode response: %v", err)
	}

	stablecoins := map[string]bool{
		"usdt": true, "usdc": true, "busd": true, "tusd": true, "dai": true, "pax": true,
	}

	cryptos := []CryptoRSI{}
	for i, coin := range data {
		if stablecoins[coin.Symbol] {
			log.Printf("[DEBUG] Skipping stablecoin: %s (rank: %d, market cap: %.0f)", coin.Symbol, i+1, coin.MarketCap)
			continue
		}

		symbol := strings.ToUpper(coin.Symbol) + "USDT"
		if _, err := fetchRSIFunc(symbol, "15m", 14); err != nil {
			log.Printf("[DEBUG] Skipping unsupported symbol: %s (rank: %d, market cap: %.0f)", symbol, i+1, coin.MarketCap)
			continue
		}

		if len(cryptos) < limit {
			cryptos = append(cryptos, CryptoRSI{
				Symbol:    symbol,
				MarketCap: coin.MarketCap,
				Rank:      i + 1, // Set market cap rank (1-based)
			})
		}
	}

	return cryptos, nil
}
