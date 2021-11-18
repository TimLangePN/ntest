package cmd

import (
	"os"

	"github.com/bschaatsbergen/hhop/pkg/https"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	Address string

	rootCmd = &cobra.Command{
		Use:   "ntest",
		Short: "ntest - here to trace network routes",
		Long:  `ntest - here to trace network routes`,
		Run: func(cmd *cobra.Command, args []string) {

			if condition := Address == ""; condition {
				logrus.Error("address flag is required")
				os.Exit(1)
			}

			PerformTests(Address)
		},
	}
)

func init() {
	configureLogLevel()

	rootCmd.Flags().StringVarP(&Address, "address", "a", "", "address to trace")
}

func configureLogLevel() {
	logLevelStr, ok := os.LookupEnv("LOG_LEVEL")
	if !ok {
		logLevelStr = "info"
	}
	logLevel, err := logrus.ParseLevel(logLevelStr)
	if err != nil {
		logLevel = logrus.InfoLevel
	}
	logrus.SetLevel(logLevel)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		logrus.Error(err)
		os.Exit(1)
	}
}

func PerformTests(Address string) {
	logrus.Info("Performing tests against: ", Address)

	// addresses := dns.LookupDnsRecords(Address)

	// for _, address := range addresses {
	// 	logrus.Info(address)
	// }

	supportsHttps := https.SupportsHttps(Address)

	if supportsHttps {
		logrus.Info("https is supported")
	}
}
