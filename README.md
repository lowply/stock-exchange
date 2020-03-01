# Stock & Exchange

Scrapes https://stocks.finance.yahoo.co.jp and http://www.murc-kawasesouba.jp and get the stock price and USD-JPY TTM of the day. Useful for your RSU calculation in Japan.

## Installation

If you have [Go](https://golang.org/) installed on your machine, just run:

```
go get github.com/lowply/stock-exchange
```

If you don't:

- Download [the latest release](https://github.com/lowply/stock-exchange/releases)
- Unarchive it
- Move the `stock-exchange` binary to `/usr/local/bin` or equivalent directory

## Usage

Run:

```
stock-exchange MSFT 2020-02-14
```

Result:

```
Querying https://stocks.finance.yahoo.co.jp/us/history/MSFT?sy=2020&sm=2&sd=14&ey=2020&em=2&ed=14&tm=d ...
2020/02/14: MSFT = 185.35
Querying http://www.murc-kawasesouba.jp/fx/past/index.php?id=200214 ...
2020/02/14: USD TTM = 109.89
```

## Deployment

```
git tag v0.x.x && git push origin v0.x.x
```

## Disclaimer

Please use this tool with your own responsibility.
