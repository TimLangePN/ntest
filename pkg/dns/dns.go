package dns

import (
	"net"
	"strings"

	"github.com/bschaatsbergen/ntest/pkg/model"
	"github.com/sirupsen/logrus"
)

// LookupHost Looks up the given host using the local resolver.
func LookupHost(options model.Options) {
	addresses, err := net.LookupHost(options.ParsedAddress)
	if err != nil {
		logrus.Error(err)
	}

	log.Infof("Hosts: %s", strings.Join(addresses[:], ", "))
}
