package cmd

import (
	"os"

	"github.com/bschaatsbergen/ntest/pkg/dns"
	"github.com/bschaatsbergen/ntest/pkg/https"
	"github.com/bschaatsbergen/ntest/pkg/model"
	"github.com/bschaatsbergen/ntest/pkg/ping"
	"github.com/bschaatsbergen/ntest/pkg/tls"
	"github.com/bschaatsbergen/ntest/pkg/utils"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	version string

	options model.Options

	rootCmd = &cobra.Command{
		Use:     "ntest",
		Short:   "ntest - cross-platform cli app that runs multiple tests against any address.",
		Long:    `ntest - cross-platform cli app that runs multiple tests against any address.`,
		Version: version, // The version is set during the build by making using of `go build -ldflags`
		Run: func(cmd *cobra.Command, args []string) {

			if condition := options.RawAddress == ""; condition {
				logrus.Error("address flag is required")
				os.Exit(1)
			}

			configureLogLevel(options.Debug)
			Test(options)
		},
	}
)

func init() {
	rootCmd.Flags().StringVarP(&options.RawAddress, "address", "a", "", "ip or address to perform tests against")
	rootCmd.Flags().IntVar(&options.PacketCount, "packet-count", 1, "amount of packets that should be sent")
	rootCmd.Flags().BoolVarP(&options.Debug, "debug", "d", false, "set log level to debug")
	rootCmd.Flags().BoolVar(&options.ReturnResponseHeaders, "headers", false, "return the response headers")
}

// configureLogLevel If an existing log level environment variable is present, re-use that to configure logrus.
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
		logrus.Fatal(err)
	}
}

// Test Runs a set of tests against the provided address.
func Test(options model.Options) {

	// We first parse the given address and return the address.Host.
	parsedAddress, err := utils.ParseAddress(options.RawAddress)

	options.ParsedAddress = parsedAddress

	if err != nil {
		logrus.Fatal(err)
	}

	ping.Ping(options)

	if options.ReturnResponseHeaders {
		https.LogResponseHeaders(options)
	}

	https.TestHttpsRedirect(options)

	tls.TestTLSCertificate(options)

	dns.LookupHost(options)
}
