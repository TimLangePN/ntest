package ping

import (
	"github.com/go-ping/ping"
	log "github.com/sirupsen/logrus"
)

var (
	Statistics *ping.Statistics // Ping statistics that are set when calling pingAddress.
)

// pingAddress Ping a host and return a set of statistics.
func pingAddress(Address string, PacketCount int) {
	pinger, err := ping.NewPinger(Address)
	if err != nil {
		log.Fatal(err)
	}

	defer pinger.Stop() // Gracefully close the pinger after it's done.

	pinger.Count = PacketCount

	if PacketCount == 1 {
		log.Debugf("Sending %d packet to: %s", pinger.Count, pinger.Addr())
	} else {
		log.Debugf("Sending %d packets to: %s", pinger.Count, pinger.Addr())
	}

	err = pinger.Run() // Blocks until finished.
	if err != nil {
		log.Fatal(err)
	}

	log.Debugf("Packets: Sent = %d, Received = %d, Lost = %d (%d%% loss)", pinger.PacketsSent, pinger.PacketsRecv, pinger.PacketsSent-pinger.PacketsRecv, int(pinger.Statistics().PacketLoss))

	Statistics = pinger.Statistics()
}

// Ping is a thin wrapper around the pingAddress function
func Ping(Address string, PacketCount int) {
	pingAddress(Address, PacketCount)

	if Statistics.PacketLoss > 0 {
		log.Warnf("Detected %d%% packet loss!", int(Statistics.PacketLoss))
	}

	log.Infof("Round-trip time: %dms", Statistics.AvgRtt.Milliseconds())
}
