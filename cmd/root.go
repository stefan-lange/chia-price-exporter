package cmd

import (
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "chia-price-exporter",
	Short: "Prometheus metric exporter for the current chia price",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	var (
		metricsPort int
		logLevel    string
	)

	rootCmd.PersistentFlags().IntVar(&metricsPort, FlagMetricsPort, 9915, "The port the metrics server binds to")
	rootCmd.PersistentFlags().StringVar(&logLevel, FlagLogLevel, "info", "How verbose the logs should be. panic, fatal, error, warn, info, debug, trace")

	err := viper.BindPFlag(FlagMetricsPort, rootCmd.PersistentFlags().Lookup(FlagMetricsPort))
	if err != nil {
		log.Fatalln(err.Error())
	}
	err = viper.BindPFlag(FlagLogLevel, rootCmd.PersistentFlags().Lookup(FlagLogLevel))
	if err != nil {
		log.Fatalln(err.Error())
	}
}
