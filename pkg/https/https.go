package https

import (
	"net/http"
	"strings"

	"github.com/sirupsen/logrus"
)

const (
	httpsProtocol = "https://"
)

func SupportsHttps(Address string) bool {
	resp, err := http.Get(Address)

	if err != nil {
		logrus.Error(err)
	}

	// We assume that a redirect has been performed.
	if strings.HasPrefix(resp.Request.URL.String(), httpsProtocol) {
		return true
	} else {
		return false
	}
}
