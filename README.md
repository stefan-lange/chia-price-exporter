# Chia Price Exporter

[![Go Report Card](https://goreportcard.com/badge/github.com/stefan-lange/chia-price-exporter)](https://goreportcard.com/report/github.com/stefan-lange/chia-price-exporter)

Chia Price Exporter is a standalone application which exports metrics to
a [Prometheus](https://github.com/prometheus/prometheus) compatible `/metrics` endpoint.

## Usage

`chia-price-exporter serve` will start the metrics exporter on the default port of `9915`. Metrics will be available
at `<hostname>:9915/metrics`.

To see further supported commands and flags try `chia-price-exporter --help`.

## Metrics

The following metrics are supported and fully compatible with the
great [Chia Monitor](https://github.com/philippnormann/chia-monitor) and the [Chia Farm Dashboard](https://github.com/stefan-lange/chia-farm-dashboard).

| Metric                   | Description |
|--------------------------|-------------|
| `chia_price_usd_cent`    | USD price   |
| `chia_price_eur_cent`    | EUR price   |
| `chia_price_btc_satoshi` | BTC price   |
| `chia_price_eth_gwei`    | ETH price   |

## Prometheus configuration

Add this job configuration to the `scrape_configs` section in your `prometheus.yml`.

```yaml
scrape_configs:
    -   job_name: 'chia-price-exporter'
        scrape_interval: 60s
        static_configs:
            -   targets: [ '<<CHIA-PRICE-EXPORTER-HOSTNAME>>:9915' ]
```

## Development (quick start)

### using go

```shell
# build and run
make build
./bin/chia-price-exporter serve
# test manually
curl -v localhost:9915/metrics
```

### using docker

```shell
# build and run
docker run --rm -p "9915:9915" -it $(docker build -q .) serve
# test manually
curl -v localhost:9915/metrics
```

play with executable

```shell
docker run --rm -it $(docker build -q .)
```

