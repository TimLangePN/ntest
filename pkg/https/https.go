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

func TestHttpsRedirect(Address string) {
	resp, err := http.Get(httpProtocol + Address)

	if err != nil {
		logrus.Error(err)
	}

	defer resp.Body.Close() // Free file descriptor to prevent resource leak.

	if strings.HasPrefix(resp.Request.URL.String(), httpsProtocol) {
		logrus.Info("HTTPS redirect detected")
	} else {
		logrus.Warn("HTTPS redirect undetected")
	}
}
