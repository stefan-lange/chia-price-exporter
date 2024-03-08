package collectors

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/prometheus/client_golang/prometheus"
	log "github.com/sirupsen/logrus"
)

var (
	vsCurrencies = []string{"BTC", "ETH", "EUR", "USD"}
	priceAPI     = fmt.Sprintf("https://api.coingecko.com/api/v3/simple/price?ids=chia&vs_currencies=%s", strings.Join(vsCurrencies, ","))
	client       = &http.Client{}
)

// CoinGeckoPriceResponse for the coingecko.com price api v3
type CoinGeckoPriceResponse struct {
	Chia struct {
		BTC float64 `json:"btc"`
		ETH float64 `json:"eth"`
		EUR float64 `json:"eur"`
		USD float64 `json:"usd"`
	} `json:"chia"`
}

// NewCoinGeckoPriceCollector to return a new instance
func NewCoinGeckoPriceCollector() *PriceCollector {
	return &PriceCollector{}
}

// Describe registers metrics
func (c *PriceCollector) Describe(ch chan<- *prometheus.Desc) {
	c.InitMetrics(ch)
}

// Collect is called by the Prometheus registry when collecting metrics.
func (c *PriceCollector) Collect(ch chan<- prometheus.Metric) {
	coinGeckoPriceResponse, err := getCurrentPrices()
	if err != nil {
		UpdateGauge(c.dataOk, 0, ch)
		log.Warn(err)
		return
	}
	UpdateGauge(c.dataOk, 1, ch)

	// update price metrics
	UpdateGauge(c.btcSatoshi, coinGeckoPriceResponse.Chia.BTC*10e7, ch)
	UpdateGauge(c.ethGwei, coinGeckoPriceResponse.Chia.ETH*10e8, ch)
	UpdateGauge(c.eurCents, coinGeckoPriceResponse.Chia.EUR*100, ch)
	UpdateGauge(c.usdCents, coinGeckoPriceResponse.Chia.USD*100, ch)
}

// GetCurrentPrices retrieve prices from API
func getCurrentPrices() (*CoinGeckoPriceResponse, error) {
	resp, err := client.Get(priceAPI)
	if err != nil {
		return nil, err
	}

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	err = resp.Body.Close()
	if err != nil {
		return nil, err
	}

	// process results
	coinGeckoPriceResponse := CoinGeckoPriceResponse{}
	log.Debugf("responseBody: %s\n", responseBody)
	err = json.Unmarshal(responseBody, &coinGeckoPriceResponse)
	if err != nil {
		return nil, err
	}

	return &coinGeckoPriceResponse, nil
}
