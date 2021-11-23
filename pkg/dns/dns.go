package dns

import (
	"net"
	"strings"

	"github.com/sirupsen/logrus"
)

// LookupHost Looks up the given host using the local resolver.
func LookupHost(Address string) {
	addresses, err := net.LookupHost(Address)
	if err != nil {
		logrus.Error(err)
	}

	logrus.Infof("DNS hosts: %s", strings.Join(addresses[:], ", "))
}
