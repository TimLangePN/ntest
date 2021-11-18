package https

import (
	"net/http"
	"strings"

	"github.com/sirupsen/logrus"
)

const (
	httpProtocol  = "http://"
	httpsProtocol = "https://"
)

func HttpsRedirectCheck(Address string) {
	resp, err := http.Get(Address)

	if err != nil {
		logrus.Error(err)
	}

	if strings.HasPrefix(resp.Request.URL.String(), httpsProtocol) {
		logrus.Info("HTTPS redirect detected")
	} else {
		logrus.Info("HTTPS redirect did *not* happen")
	}
}
