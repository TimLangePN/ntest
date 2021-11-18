package tls

import (
	"crypto/tls"
	"math"
	"os"
	"time"

	"github.com/sirupsen/logrus"
)

const (
	httpsPort = ":443"
)

// Performs a TLS handshake with the given server and returns the TLS connection
// in order to be able to perform further operations on it, e.g. calculate the leaf certificate's expiration date.
func TestTLSCertificate(Address string) {

	conn, err := tls.Dial("tcp", Address+httpsPort, nil)

	if err != nil {
		logrus.Error(err)
		os.Exit(1)
	}

	// Check if the peer certificate chain is valid.
	err = conn.VerifyHostname(Address)

	if err != nil {
		logrus.Error(err)
	}

	// Get end-entity certificate (leaf certificate) from the X.509 certificate chain.
	leafCert := conn.ConnectionState().PeerCertificates[0]

	// Calculate the time left until the certificate expires.
	// This is rather complex, we can split it up later.
	daysUntilExpiration := math.Round(float64(time.Until(leafCert.NotAfter).Hours()) / 24)

	// Print the certificate expiration in number of days left.
	if daysUntilExpiration > 0 {
		logrus.Infof("Certificate for %s expires in %v days", leafCert.DNSNames, daysUntilExpiration)

	} else {
		logrus.Errorf("Certificate for %s is expired!", Address)
	}
}
