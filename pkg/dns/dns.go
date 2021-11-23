package dns

import (
	"net"
	"strings"

	log "github.com/sirupsen/logrus"
)

// LookupHost Looks up the given host using the local resolver.
func LookupHost(Address string) {
	addresses, err := net.LookupHost(Address)
	if err != nil {
		log.Error(err)
	}

	log.Infof("DNS hosts: %s", strings.Join(addresses[:], ", "))
}
