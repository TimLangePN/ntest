package latency

import (
	"os"

	"github.com/go-ping/ping"
	"github.com/sirupsen/logrus"
)

// Ping function to ping a host and return a set of statistics.
// Uses the github.com/go-ping/ping module.
func Ping(Address string) *ping.Statistics {

	pinger, err := ping.NewPinger(Address)
	if err != nil {
		logrus.Error(err)
		os.Exit(1)
	}

	defer pinger.Stop() // Gracefully close the pinger after it's done.

	pinger.Count = 1

	logrus.Debugf("Sending %d packet to: %s", pinger.Count, pinger.Addr())

	err = pinger.Run() // Blocks until finished.
	if err != nil {
		logrus.Error(err)
		os.Exit(1)
	}

	logrus.Debugf("Packets: Sent = %d, Received = %d, Lost = %d (%d%% loss)", pinger.PacketsSent, pinger.PacketsRecv, pinger.PacketsSent-pinger.PacketsRecv, int(pinger.Statistics().PacketLoss))

	return pinger.Statistics() // get send/receive/duplicate/rtt stats
}

// Exposes the average round trip time in milliseconds resolved from the Ping function.
func MeasureLatency(Address string) {

	stats := Ping(Address)

	logrus.Infof("Round-trip time: %dms", stats.AvgRtt.Milliseconds()) // Fix this later, so it returns a float (e.g. `13.4123ms`).
}
