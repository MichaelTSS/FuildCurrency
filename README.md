# FluidCurrency ![Language Badge](https://img.shields.io/badge/Language-Go-blue.svg) ![License Badge](https://img.shields.io/badge/License-MIT-blue.svg) ![Status Badge](https://img.shields.io/badge/Status-Alpha-yellow.svg)

FluidCurrency is a Go utility to fetch currencies, crypto-currencies and precious metal exchange rates. It does come with a number of providers (see below) for maximum flexibility. This library provides a minimalist API and a default provider. Some developers need a high frequency rate, others only a limited of currencies. Check the providers comparison table and pick the one that bests suits you.

## Project Status

GoCurrency is on alpha. Pull Requests [are welcome](CONTRIBUTING.md)

## Features

- Open source - You can check out our code
- Easy to use
- Come with different providers to avoid vendor lock-in

## Providers

Provider                  | History | Frequency Rate       | Streaming | SSL | Uptime (advertised) | Response time (advertised) | Plan                                            | Currencies | Crypto-currencies | Precious Metals
------------------------- | ------- | -------------------- | --------- | --- | ------------------- | -------------------------- | ----------------------------------------------- | ---------- | ----------------- | ---------------
**live-rates.com**        | ---     | 1/sec                | ---       | --- | ---                 | ---                        | free                                            | ---        | ---               | ---
**European Central Bank** | 1999    | 1/day at 3PM CET     | ---       | --- | ---                 | ---                        | free                                            | limited    | NO                | NO
**Fixer.io**              | 1999    | 1/day at 4PM CET     | ---       | YES | 99.98%              | ~5ms                       | free (no API key, auto throttles)               | limited    | NO                | NO
**XE**                    | NO      | real-time            | ---       | --- | ---                 | ---                        | from $799/yr (no free trial)                    | > 190      | NO                | ---
**Blockchain**            | ---     | ---                  | ---       | --- | ---                 | ---                        | free (optional API key for a higher rate limit) | ---        | YES               | NO
**Coin Desk**             | ---     | 10/sec               | ---       | --- | ---                 | YES                        | free                                            | ---        | YES               | NO
**Currency Layer**        | ---     | from 1/min to 1/hour | ---       | --- | ---                 | ---                        | free plan (with API key) then $95/yr            | > 160      | NO                | YES
**Xiniche**               | 1999    | real-time            | ---       | --- | ---                 | ---                        | 7-day free trial then $3k/yr                    | ---        | NO                | ---
**Oanda**                 | 1990    | real-time            | ---       | --- | ---                 | ---                        | 30-day free trial then $4.5k/yr                 | ---        | NO                | ---

## Installation

```bash
$ go get github.com/MichaelTSS/FuildCurrency
```

## TODO

- Providers table
  - Fill holes

- Code

  - Add Rates interface
  - Add Rates request parameters

## License

FuildCurrency is [MIT licenced](LICENCE).
