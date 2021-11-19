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

Foobar

## Installation

### Binaries

If you prefer grabbing ntest its binaries, download the latest from the the **[GitHub releases](https://github.com/bschaatsbergen/ntest/releases)** page.

### Brew

```sh
❯ brew tap bschaatsbergen/ntest
❯ brew install ntest
```

## Usage

Foobar

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

ntest currently performs the following tests:

```sh
❯ ntest -a bschaatsbergen.com
INFO[0020] Round-trip time: 13ms                        
INFO[0036] HTTPS redirect detected                      
INFO[0036] Certificate for bschaatsbergen.com, *.bschaatsbergen.com expires in 186 days 
INFO[0036] DNS hosts: 52.222.138.52, 52.222.138.100, 52.222.138.23, 52.222.138.36
```

## Contributing

If you have a suggestion that would make this better, please fork the repo and create a pull request. You can also simply open an issue with the tag "enhancement".

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
