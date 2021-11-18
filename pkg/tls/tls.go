package tls

import (
	"crypto/tls"
	"fmt"
	"math"
	"os"
	"time"

	"github.com/sirupsen/logrus"
)

const (
	httpsPort = ":443"
)

func TlsCertificateCheck(Address string) {

	conn, err := tls.Dial("tcp", Address+httpsPort, nil)

	if err != nil {
		logrus.Error(err)
		os.Exit(1)
	}

	err = conn.VerifyHostname(Address)

	if err != nil {
		logrus.Error(err)
		os.Exit(1)
	}

	expirationDate := conn.ConnectionState().PeerCertificates[0].NotAfter

	daysUntilExpiration := math.Round(float64(time.Until(expirationDate).Hours()) / 24)

	if daysUntilExpiration > 1 {
		logrus.Info(fmt.Sprintf("Certificate for %s expires in %v days", Address, daysUntilExpiration))
	} else {
		logrus.Error(fmt.Sprintf("Certificate for %s is expired!", Address))
	}
}
