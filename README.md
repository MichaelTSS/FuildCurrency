# ![FluidCurrency](http://image.prntscr.com/image/e1498265eef040bc956e388c35d8f93a.png)

# FluidCurrency ![Language Badge](https://img.shields.io/badge/Language-Go-blue.svg) ![License Badge](https://img.shields.io/badge/License-MIT-blue.svg) ![Status Badge](https://img.shields.io/badge/Status-Alpha-yellow.svg)

FluidCurrency is a Go utility to fetch currencies, crypto-currencies and precious metal exchange rates. It does come with a number of providers (see below) for maximum flexibility. This library provides a minimalist API and a default provider. Some developers need a high frequency rate, others only a limited of currencies. Check the providers comparison table and pick the one that bests suits you.

**Table of Contents**

- [Demo](#demo)
- [Project Status](#project-status)
- [Features](#features)
- [Installation](#installation)
- [Contributing](#contributing)
  - [Bug Reports & Feature Requests](#bug-reports--feature-requests)
  - [Developing](#developing)
- [License](#license)

## Demo

[![asciicast](http://image.prntscr.com/image/2f33d4153f794d15bd95d2d533adab98.png)](https://asciinema.org/a/107878?t=10)

## Project Status

GoCurrency is on alpha. Pull Requests [are welcome](CONTRIBUTING.md)

## Features

- Open source - You can check out our code
- Secure
- Always up-to-date
- Use decimal type
- 100% satisfaction guaranteed
- It's perfect to convert amounts between currencies
- STUPIDLY [EASY TO USE](https://github.com/DiSiqueira/GoCurrency#usage)
- Very fast start up and response time
- Uses native libs

## Installation

### Go Get

```bash
$ go get github.com/MichaelTSS/FuildCurrency
```

## Usage

### Get all available currencies

```golang
package main

import (
	"fmt"

	"github.com/MichaelTSS/FuildCurrency"
)

func main() {
	curList, _ := gocurrency.AvailableCurrencies()

	for _, currency := range curList {
		fmt.Println(currency.Description)
	}
}
```

Output:
```
SEK Sweden, kronor
ATS Austria, shilling
AUD Australian, dollar
BEF Belgien, franc
BRL Brazilien, real
CAD Canada, dollar
CHF Switzerland, francs
CNY China, yuan renminbi
CYP Cyprus, pound
CZK Czech Republic, koruna
DEM Germany, mark
DKK Denmark, krone
EEK Estonian, kroon
ESP Spain, pesetas
EUR Euroland, euro
FIM Finland, marka
FRF France, franc
GBP Great Britain, pound
GRD Greece, drachmer
HKD Hong Kong, dollar
HUF Hungary, forint
IDR Indonesia, rupiah
IEP Ireland, pund
INR India, rupee
ISK Iceland, kronor
ITL Italy, lire
JPY Japan, yen
KRW South Korea, won
KWD Kuwait, dinar
LTL Lithuania,  litas
LVL Latvia, lat
MAD Morocko, dirham
MXN Mexico, nuevo peso
MYR Malaysia, ringgit
NLG Dutchland, guilder
NOK Norway, krone
NZD New Zealand, dollar
PLN Poland, zloty
PTE Portugal, escudo
RUB Russia, rouble
SAR Saudi Arabia, riyal
SGD Singapore, dollar
SIT Slovenia, tolar
SKK Slovakia, koruna
THB Thailand, baht
TRL Turkey, lira
TRY Turkey, new lira
USD US, dollar
ZAR South Africa, rand
```

## Contributing

### Bug Reports & Feature Requests

Please use the [issue tracker](https://github.com/MichaelTSS/FuildCurrency/issues) to report any bugs or file feature requests.

## License

The MIT License (MIT)

Copyright (c) 2013-2017 Michel Tresseras

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.  IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
