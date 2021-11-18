package utils

import (
	"fmt"
	"net/url"
	"strings"
)

// Checks whether the given string is a valid IPv4 address.
func IsIPv4(address string) bool {
	return strings.Count(address, ":") < 2
}

// Checks whether the given string is a valid IPv6 address.
func IsIPv6(address string) bool {
	return strings.Count(address, ":") >= 2
}

// Returns a local listener address for IPv4.
func ReturnIPv4LocalListener() string {
	return "0.0.0.0:0"
}

// Returns a local listener address for IPv6.
func ReturnIPv6LocalListener() string {
	return "::"
}

// Wrapper for url.ParseRequestURI(), handles the case where the `Host` property might end up empty.
func ParseAddress(rawurl string) (domain string, err error) {
	u, err := url.ParseRequestURI(rawurl)
	if err != nil || u.Host == "" {
		u, repErr := url.ParseRequestURI("https://" + rawurl)
		if repErr != nil {
			fmt.Printf("Could not parse raw url: %s, error: %v", rawurl, err)
			return
		}
		domain = u.Host
		err = nil
		return
	}

	domain = u.Host
	return
}
