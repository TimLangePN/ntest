package https

import (
	"net/http"
	"strings"

	"github.com/bschaatsbergen/ntest/pkg/model"
	"github.com/sirupsen/logrus"
	"github.com/xlab/treeprint"
)

const (
	httpProtocol  = "http://"
	httpsProtocol = "https://"
)

// TestHttpsRedirect Tests whether a host redirects HTTP to HTTPS.
func TestHttpsRedirect(options model.Options) {
	resp, err := http.Get(httpProtocol + options.ParsedAddress)

	if err != nil {
		logrus.Error(err)
	}

	defer resp.Body.Close() // Free file descriptor to prevent resource leak.

	if strings.HasPrefix(resp.Request.URL.String(), httpsProtocol) {
		logrus.Info("HTTPS redirect detected")
	} else {
		logrus.Warn("HTTPS redirect undetected")
	}

	logrus.Debugf("HTTPS redirect returned: %s", resp.Status)
}

// LogResponseHeaders does a HTTP call and logs the repsonse headers in a tree-like format
func LogResponseHeaders(options model.Options) {
	headers := getResponseHeaders(options.ParsedAddress)
	prettyLogHeaders(headers)
}

func getResponseHeaders(address string) http.Header {
	resp, err := http.Get(httpsProtocol + address)

	if err != nil {
		logrus.Error(err)
	}

	defer resp.Body.Close() // Free file descriptor to prevent resource leak.

	if err != nil {
		logrus.Error(err)
	}

	logrus.Debugf("Received %v response headers", len(resp.Header))

	return resp.Header
}

func prettyLogHeaders(m map[string][]string) {
	tree := treeprint.NewWithRoot("Response headers:")

	for i, v := range m {
		tree.AddMetaBranch(i, v) // Add a meta branch for each header.
	}

	logrus.Info(tree.String())
}
