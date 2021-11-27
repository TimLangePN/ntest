package ping

import (
	"github.com/bschaatsbergen/ntest/pkg/model"
	"github.com/go-ping/ping"
	"github.com/sirupsen/logrus"
)

// Ping is a thin wrapper around the pingAddress function
func Ping(options model.Options) {
	stats := pingAddress(options.ParsedAddress, options.PacketCount)

	if stats.PacketLoss > 0 {
		logrus.Warnf("Detected %d%% packet loss!", int(stats.PacketLoss))
	}

	logrus.Infof("Round-trip time: %dms", stats.AvgRtt.Milliseconds())
}

// pingAddress Ping a host and return a set of statistics.
func pingAddress(address string, packetCount int) *ping.Statistics {
	pinger, err := ping.NewPinger(address)
	if err != nil {
		logrus.Fatal(err)
	}

	defer pinger.Stop() // Gracefully close the pinger after it's done.

	pinger.Count = packetCount

	if packetCount == 1 {
		logrus.Debugf("Sending %d packet to: %s", pinger.Count, pinger.Addr())
	} else {
		logrus.Debugf("Sending %d packets to: %s", pinger.Count, pinger.Addr())
	}

	err = pinger.Run() // Blocks until finished.
	if err != nil {
		logrus.Fatal(err)
	}

	logrus.Debugf("Packets: Sent = %d, Received = %d, Lost = %d (%d%% loss)", pinger.PacketsSent, pinger.PacketsRecv, pinger.PacketsSent-pinger.PacketsRecv, int(pinger.Statistics().PacketLoss))

	return pinger.Statistics()
}
