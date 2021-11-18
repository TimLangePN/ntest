package utils

import "strings"

func IsIPv4(address string) bool {
	return strings.Count(address, ":") < 2
}

func IsIPv6(address string) bool {
	return strings.Count(address, ":") >= 2
}

func ReturnIPv4LocalListener() string {
	return "0.0.0.0:0"
}

func ReturnIPv6LocalListener() string {
	return "::"
}
