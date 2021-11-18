package dns

import (
	"net"

	"github.com/sirupsen/logrus"
)

// Looks up the given host using the local resolver.
func LookupHost(Address string) {
	addresses, err := net.LookupHost(Address)
	if err != nil {
		logrus.Error(err)
	}

	logrus.Infof("Hosts: %s", addresses)
}
