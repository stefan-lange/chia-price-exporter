package collectors

import "github.com/prometheus/client_golang/prometheus"

const (
	//MetricNamespace the namespace of the exported metrics
	MetricNamespace = "chia"

	//MetricSubsystem the subsystem / sub category of the exported metrics
	MetricSubsystem = "price"
)

// PriceCollector internal data structure of the price collector
type PriceCollector struct {
	prometheus.Collector

	// status metric (maybe this can be removed in the future)
	dataOk prometheus.Gauge

	// price metrics
	btcSatoshi prometheus.Gauge
	ethGwei    prometheus.Gauge
	eurCents   prometheus.Gauge
	usdCents   prometheus.Gauge
}

// InitMetrics registers any metrics (gauges, counters, etc). Use this in the 'Describe()' of the current collector implementation
func (c *PriceCollector) InitMetrics(ch chan<- *prometheus.Desc) {

	c.dataOk = NewGauge("data_ok", "Was the last API scraping ok", ch)

	c.btcSatoshi = NewGauge("btc_satoshi", "Current Chia price in BTC satoshi", ch)
	c.ethGwei = NewGauge("eth_gwei", "Current Chia price in ETH gwei", ch)
	c.eurCents = NewGauge("eur_cent", "Current Chia price in EUR cent", ch)
	c.usdCents = NewGauge("usd_cent", "Current Chia price in USD cent", ch)
}

// NewGauge helper to create and register a prometheus.Gauge
func NewGauge(name string, help string, ch chan<- *prometheus.Desc) prometheus.Gauge {
	gauge := prometheus.NewGauge(prometheus.GaugeOpts{
		Namespace: MetricNamespace,
		Subsystem: MetricSubsystem,
		Name:      name,
		Help:      help,
	})

	ch <- gauge.Desc()

	return gauge
}

// UpdateGauge helper to update a prometheus.Gauge
func UpdateGauge(gauge prometheus.Gauge, value float64, ch chan<- prometheus.Metric) {
	gauge.Set(value)
	ch <- gauge
}
