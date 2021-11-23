package model

// Options Use this to bind cobra flags, initialized in `root.go`.
type Options struct {
	Address     string
	Debug       bool
	PacketCount int
}
