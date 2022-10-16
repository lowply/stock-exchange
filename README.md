# Stock & Exchange

Scrapes https://finance.yahoo.co.jp and http://www.murc-kawasesouba.jp and get the stock price and USD-JPY TTM of the day. Useful for your RSU calculation in Japan.

## Installation

If you have [Go](https://golang.org/) installed on your machine, just run:

```
go install github.com/lowply/stock-exchange@latest
```

If you don't:

- Download [the latest release](https://github.com/lowply/stock-exchange/releases)
- Unarchive it
- Move the `stock-exchange` binary to `/usr/local/bin` or equivalent directory

## Usage

Run:

```console
stock-exchange MSFT 2020-02-14
```

Result:

```console
Querying https://finance.yahoo.co.jp/quote/MSFT/history?from=20200214&to=20200214&timeFrame=d&page=1 ...
2020/02/14: MSFT = 185.35
Querying http://www.murc-kawasesouba.jp/fx/past/index.php?id=200214 ...
2020/02/14: JPY/USD TTM = 109.89
```

## Deployment

```
git tag v0.x.x && git push origin v0.x.x
```

## Disclaimer

Please use this tool with your own responsibility.
