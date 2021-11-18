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
func CheckTLSCertificate(Address string) {

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

	// Get the `NotAfter` time for the leaf certificate
	// that the connection is verified against.
	expirationDate := conn.ConnectionState().PeerCertificates[0].NotAfter

	// Calculate the time left until the certificate expires.
	daysUntilExpiration := math.Round(float64(time.Until(expirationDate).Hours()) / 24)

	// Print the certificate expiration in number of days left.
	if daysUntilExpiration > 0 {
		logrus.WithFields(logrus.Fields{
			"days": daysUntilExpiration,
		}).Infof("Certificate for %s expires in", Address)

	} else {
		logrus.Errorf("Certificate for %s is expired!", Address)
	}
}
