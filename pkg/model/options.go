package model

// Options Use this to bind cobra flags, initialized in `root.go`.
type Options struct {
	RawAddress            string
	ParsedAddress         string
	Debug                 bool
	PacketCount           int
	ReturnResponseHeaders bool
}
