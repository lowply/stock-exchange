# Stock & Exchange

Scrapes https://stocks.finance.yahoo.co.jp and http://www.murc-kawasesouba.jp and get the stock price and USD-JPY TTM of the day. Useful for your MSFT RSU calculation in Japan.

## Installation

1. [Install latest go](https://golang.org/dl) if you haven't
2. Run `git clone https://github.com/lowply/stock-exchange.git && cd stock-exchange && make`

## Usage

Run:

```
./stockexchange MSFT 2020-02-04
```

Result:

```
Querying https://stocks.finance.yahoo.co.jp/us/history/MSFT?sy=2020&sm=2&sd=4&ey=2020&em=2&ed=4&tm=d ...
Querying http://www.murc-kawasesouba.jp/fx/past/index.php?id=200204 ...
2020年2月4日, 180.12
February 4, 2020, US Dollar, 108.61
```

## Disclaimer

Please use this tool with your own responsibility.