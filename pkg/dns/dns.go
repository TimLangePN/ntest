package dns

import (
	"net"

	"github.com/sirupsen/logrus"
)

func LookupDnsRecords(Address string) []string {
	adresses, err := net.LookupHost(Address)
	if err != nil {
		logrus.Error(err)
	}

	return adresses
}
