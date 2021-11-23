package https

import (
	"net/http"
	"strings"

	log "github.com/sirupsen/logrus"
)

const (
	httpProtocol  = "http://"
	httpsProtocol = "https://"
)

// TestHttpsRedirect Tests whether a host redirects us to https.
func TestHttpsRedirect(Address string) {
	resp, err := http.Get(httpProtocol + Address)

	if err != nil {
		log.Error(err)
	}

	defer resp.Body.Close() // Free file descriptor to prevent resource leak.

	if strings.HasPrefix(resp.Request.URL.String(), httpsProtocol) {
		log.Info("HTTPS redirect detected")
	} else {
		log.Warn("HTTPS redirect undetected")
	}

	log.Debugf("HTTPS redirect returned a: %s", resp.Status)
}
