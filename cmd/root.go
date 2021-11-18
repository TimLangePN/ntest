package cmd

import (
	"os"

	"github.com/bschaatsbergen/ntest/pkg/https"
	"github.com/bschaatsbergen/ntest/pkg/model"
	"github.com/bschaatsbergen/ntest/pkg/tls"
	"github.com/bschaatsbergen/ntest/pkg/utils"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	options model.Options

	rootCmd = &cobra.Command{
		Use:   "ntest",
		Short: "ntest - run multiple tests against any ip or address 🩺",
		Long:  `ntest - run multiple tests against any ip or address 🩺`,
		Run: func(cmd *cobra.Command, args []string) {

			if condition := options.Address == ""; condition {
				logrus.Error("address flag is required")
				os.Exit(1)
			}

			configureLogLevel(options.Debug)
			Test(options)
		},
	}
)

func init() {
	rootCmd.Flags().StringVarP(&options.Address, "address", "a", "", "ip or address to perform tests against")
	rootCmd.Flags().BoolVarP(&options.Debug, "debug", "d", false, "set log level to debug")
}

// If an existing log level environment variable is present, re-use that to configure logrus.
func configureLogLevel(debugLogsEnabled bool) {
	logLevelStr, ok := os.LookupEnv("LOG_LEVEL")
	if !ok {
		logLevelStr = "info"
	}
	if debugLogsEnabled {
		logLevelStr = "debug"
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

// Runs a set of tests against the provided address.
func Test(options model.Options) {

	// We first parse the given address and return the address.Host.
	domain, err := utils.ParseAddress(options.Address)

	if err != nil {
		logrus.Error(err)
		os.Exit(1)
	}

	// addresses := dns.LookupDnsRecords(domain)

	// for _, address := range addresses {
	// 	logrus.Info(address)
	// }

	https.HttpsRedirectCheck(domain)

	tls.CheckTLSCertificate(domain)
}
