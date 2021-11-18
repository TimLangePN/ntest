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
	resp, err := http.Get(httpProtocol + Address)

	if err != nil {
		logrus.Error(err)
	}

	if strings.HasPrefix(resp.Request.URL.String(), httpsProtocol) {
		logrus.Infof("HTTPS redirect detected, got %s", resp.Request.Response.Status)
	} else {
		logrus.Warnf("HTTPS redirect undetected got %s", resp.Request.Response.Status)
	}
}
