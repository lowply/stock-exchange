# Stock & Exchange

Scrapes https://stocks.finance.yahoo.co.jp and http://www.murc-kawasesouba.jp and get the stock price and USD-JPY TTM of the day. Useful for your RSU calculation in Japan.

## Installation

Download [the latest release](/lowply/stock-exchange/releases) and unzip it.

## Usage

Run:

```
./stock-exchange MSFT 2020-02-14
```

Result:

```
Querying https://stocks.finance.yahoo.co.jp/us/history/MSFT?sy=2020&sm=2&sd=14&ey=2020&em=2&ed=14&tm=d ...
2020/02/14: MSFT = 185.35
Querying http://www.murc-kawasesouba.jp/fx/past/index.php?id=200214 ...
2020/02/14: USD TTM = 109.89
```

## Disclaimer

Please use this tool with your own responsibility.
