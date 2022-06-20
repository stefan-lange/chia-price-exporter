package cmd

import (
	"fmt"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/stefan-lange/chia-price-exporter/collectors"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Starts the metrics server",
	Run: func(cmd *cobra.Command, args []string) {
		// read command line args
		logLevel, err := log.ParseLevel(viper.GetString(FlagLogLevel))
		if err != nil {
			log.Fatalf("Error parsing log level: %s\n", err.Error())
		}
		metricsPort := uint16(viper.GetInt(FlagMetricsPort))

		// configure logger
		log.SetLevel(logLevel)

		// holds a custom prometheus registry so that only our metrics are exported, and not the default go metrics
		registry := prometheus.NewRegistry()

		// register collectors
		registry.MustRegister(collectors.NewCoinGeckoPriceCollector())

		// let's go
		log.Fatalln(StartServer(metricsPort, registry))
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
}

// StartServer starts the metrics server
func StartServer(metricsPort uint16, prometheusRegistry *prometheus.Registry) error {
	log.Printf("Starting metrics server on port %d", metricsPort)

	http.Handle("/metrics", promhttp.HandlerFor(prometheusRegistry, promhttp.HandlerOpts{}))
	return http.ListenAndServe(fmt.Sprintf(":%d", metricsPort), nil)
}
