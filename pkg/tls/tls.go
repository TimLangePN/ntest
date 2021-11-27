package tls

import (
	"crypto/tls"
	"math"
	"strings"
	"time"

	"github.com/bschaatsbergen/ntest/pkg/model"
	"github.com/sirupsen/logrus"
)

const (
	DefaultHTTPSPort = ":443"
)

// TestTLSCertificate Performs a TLS handshake with the given server and returns the TLS connection
// in order to be able to perform further operations on it, e.g. calculate the leaf certificate's expiration date.
func TestTLSCertificate(options model.Options) {

	conn, err := tls.Dial("tcp", options.ParsedAddress+DefaultHTTPSPort, nil)

	if err != nil {
		logrus.Fatal(err)
	}

	defer conn.Close() // close the underlying TLS connection.

	// Check if the peer certificate chain is valid.
	err = conn.VerifyHostname(options.ParsedAddress)

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
		logrus.Infof("Certificate for %s expires in %v days", strings.Join(leafCert.DNSNames[:], ", "), daysUntilExpiration)

	} else {
		logrus.Errorf("Certificate for %s is expired!", options.ParsedAddress)
	}
}
