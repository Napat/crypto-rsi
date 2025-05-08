# Crypto RSI

## Installation

```sh
go install github.com/Napat/crypto-rsi@latest
```

## Usage

```sh
$ crypto-rsi

2025/05/09 00:24:42 [DEBUG] Skipping stablecoin: usdt (rank: 3, market cap: 149528794114)
2025/05/09 00:24:43 [DEBUG] Skipping stablecoin: usdc (rank: 7, market cap: 60880459200)
2025/05/09 00:24:44 [DEBUG] Skipping unsupported symbol: STETHUSDT (rank: 11, market cap: 18771289081)
2025/05/09 00:24:45 [DEBUG] Skipping unsupported symbol: WSTETHUSDT (rank: 17, market cap: 8528191424)
2025/05/09 00:24:46 [DEBUG] Skipping unsupported symbol: LEOUSDT (rank: 20, market cap: 8136523922)
+----+----------+-------+-------------------+-----------------+
| #  |  SYMBOL  |  RSI  | MARKET CAP (USD)  | MARKET CAP RANK |
+----+----------+-------+-------------------+-----------------+
|  1 | TRXUSDT  | 80.45 |    24,254,188,724 |              10 |
|  2 | SUIUSDT  | 77.06 |    13,184,324,528 |              12 |
|  3 | ETHUSDT  | 75.77 |   247,218,446,120 |               2 |
|  4 | WBTCUSDT | 73.83 |    13,052,006,288 |              13 |
|  5 | DOGEUSDT | 72.05 |    28,422,868,662 |               8 |
|  6 | BTCUSDT  | 72.04 | 2,012,170,699,716 |               1 |
|  7 | SOLUSDT  | 65.89 |    82,961,604,213 |               6 |
|  8 | BNBUSDT  | 64.36 |    90,584,085,802 |               5 |
|  9 | XRPUSDT  | 59.02 |   131,439,824,105 |               4 |
| 10 | ADAUSDT  | 58.84 |    26,314,535,833 |               9 |
+----+----------+-------+-------------------+-----------------+
```
