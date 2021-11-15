package trace

import (
	"net"

	"github.com/sirupsen/logrus"

	"github.com/bschaatsbergen/hhop/pkg/utils"
)

func Trace(Address string) {

	// Listen for incoming connections on ICMP socket.
	icmp_sock, err := net.ListenPacket("ip4:icmp", utils.ReturnIPv4LocalListener())

	if err != nil {
		logrus.Panic("Could not set a listening ICMP socket: %s\n", err)
	}

	defer icmp_sock.Close()
}
