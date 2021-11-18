package latency

import (
	"github.com/go-ping/ping"
	"github.com/sirupsen/logrus"
)

// Ping function to ping a host and return a set of statistics.
// Uses the github.com/go-ping/ping module.
func Ping(Address string) *ping.Statistics {

	pinger, err := ping.NewPinger(Address)
	if err != nil {
		logrus.Panic(err)
	}
	pinger.Count = 1
	err = pinger.Run() // Blocks until finished.
	if err != nil {
		logrus.Panic(err)
	}

	return pinger.Statistics() // get send/receive/duplicate/rtt stats
}

// Exposes the average round trip time in milliseconds resolved from the Ping function.
func MeasureLatency(Address string) {

	stats := Ping(Address)

	rttInMs := stats.AvgRtt.Microseconds() // Fix this later, so it returns a float (e.g. `13.4123ms`).

	logrus.Infof("Round-trip time: %fms", rttInMs)
}
