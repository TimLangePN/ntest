package cmd

import (
	"os"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "hhop",
	Short: "hhop - here to trace network routes",
	Long:  `hhop - here to trace network routes`,
	Run: func(cmd *cobra.Command, args []string) {
		logrus.Error("Missed or unknown command.\n\n")
		cmd.Help()
	},
}

func init() {
	configureLogLevel()
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
