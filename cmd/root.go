package cmd

import (
	"os"

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

			Test(options)
		},
	}
)

func init() {
	configureLogLevel()

	rootCmd.Flags().StringVarP(&options.Address, "address", "a", "", "ip or address to perform tests against")
	rootCmd.PersistentFlags().BoolVarP(&options.Debug, "debug", "d", false, "set log level to debug")
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

func Test(options model.Options) {

	domain, err := utils.ParseAddress(options.Address)

	if err != nil {
		logrus.Error(err)
		os.Exit(1)
	}

	// addresses := dns.LookupDnsRecords(domain)

	// for _, address := range addresses {
	// 	logrus.Info(address)
	// }

	logrus.Info(domain)

	// https.HttpsRedirectCheck(domain)

	tls.TlsCertificateCheck(domain)

	logrus.Info("done")
}
