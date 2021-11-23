package cmd

import (
	"os"

	"github.com/bschaatsbergen/ntest/pkg/dns"
	"github.com/bschaatsbergen/ntest/pkg/https"
	"github.com/bschaatsbergen/ntest/pkg/model"
	ping "github.com/bschaatsbergen/ntest/pkg/ping"
	"github.com/bschaatsbergen/ntest/pkg/tls"
	"github.com/bschaatsbergen/ntest/pkg/utils"
	log "github.com/sirupsen/logrus"
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

			if condition := options.Address == ""; condition {
				log.Error("address flag is required")
				os.Exit(1)
			}

			configureLogLevel(options.Debug)
			Test(options)
		},
	}
)

func init() {
	rootCmd.Flags().StringVarP(&options.Address, "address", "a", "", "ip or address to perform tests against")
	rootCmd.Flags().IntVar(&options.PacketCount, "packet-count", 1, "amount of packets that should be sent")
	rootCmd.Flags().BoolVarP(&options.Debug, "debug", "d", false, "set log level to debug")
}

// configureLogLevel If an existing log level environment variable is present, re-use that to configure log.
func configureLogLevel(debugLogsEnabled bool) {
	logLevelStr, ok := os.LookupEnv("LOG_LEVEL")
	if !ok {
		logLevelStr = "info"
	}
	if debugLogsEnabled {
		logLevelStr = "debug"
	}
	logLevel, err := log.ParseLevel(logLevelStr)
	if err != nil {
		logLevel = log.InfoLevel
	}
	log.SetLevel(logLevel)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

// Test Runs a set of tests against the provided address.
func Test(options model.Options) {

	// We first parse the given address and return the address.Host.
	domain, err := utils.ParseAddress(options.Address)

	if err != nil {
		log.Fatal(err)
	}

	ping.MeasureLatency(domain, options.PacketCount)

	https.TestHttpsRedirect(domain)

	tls.TestTLSCertificate(domain)

	dns.LookupHost(domain)
}
