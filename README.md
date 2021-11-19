<p align="center"><img src="https://twemoji.maxcdn.com/2/svg/1fa7a.svg" height="64" alt="Project Logo"></p>
<h3 align="center">ntest</h3>
<p align="center">ntest is a cross-platform cli app that runs multiple tests against any address.</p>
<p align="center">
    <a href="https://github.com/bschaatsbergen/ntest/releases"><img src="https://img.shields.io/github/downloads/bschaatsbergen/ntest/total.svg" alt="GitHub Downloads"></a>
    <a href="https://github.com/bschaatsbergen/ntest/releases/latest"><img src="https://img.shields.io/github/release/bschaatsbergen/ntest.svg" alt="Latest Release"></a>
    <a href="https://github.com/bschaatsbergen/ntest/actions/workflows/go-ci.yaml"><img src="https://img.shields.io/github/workflow/status/bschaatsbergen/ntest/Go" alt="Build Status"></a>
    <a href="https://github.com/bschaatsbergen/ntest/issues"><img src="https://img.shields.io/badge/contributions-welcome-ff69b4.svg" alt="Contributions Welcome"></a>
</p>

## About ntest

Having the ability to run common tests against any domain often comes in quite handy. `ntest` aims to provide you with the most necessary tests you would want to run against a domain.

## Installation

### Binaries

If you prefer grabbing ntest its binaries, download the latest from the the **[GitHub releases](https://github.com/bschaatsbergen/ntest/releases)** page.

### Brew

```sh
❯ brew tap bschaatsbergen/ntest
❯ brew install ntest
```

## Usage

Having the ability to run common tests against any domain often comes in quite handy. `ntest` aims to provide you with the most necessary tests you would want to run against a domain.

```sh
❯ ntest -h
ntest - run multiple tests against any ip or address 🩺

Usage:
  ntest [flags]

Flags:
  -a, --address string     ip or address to perform tests against
  -d, --debug              set log level to debug
  -h, --help               help for ntest
      --packet-count int   amount of packets that should be sent (default 1)
```

`ntest` currently performs the following tests:

```sh
❯ ntest -a bschaatsbergen.com
INFO[0020] Round-trip time: 13ms
INFO[0036] HTTPS redirect detected
INFO[0036] Certificate for bschaatsbergen.com, *.bschaatsbergen.com expires in 186 days
INFO[0036] DNS hosts: 52.222.138.52, 52.222.138.100, 52.222.138.23, 52.222.138.36
```

It's also possible measure the average round-trip time by sending multiple packets.
Additionally, adding the `-d` flag allows you to see the debug logs regarding the packets.

```sh
❯ ntest -a bschaatsbergen.com --packet-count 5 -d
DEBU[0000] Parsed address: bschaatsbergen.com
DEBU[0000] Sending 5 packets to: bschaatsbergen.com
DEBU[0004] Packets: Sent = 5, Received = 5, Lost = 0 (0% loss)
INFO[0004] Round-trip time: 14ms
...
```

## Contributing

If you have a suggestion that would make `ntest` better, please fork the repo and create a pull request. You can also simply open an issue with the tag "enhancement".

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
