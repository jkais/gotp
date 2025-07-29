# GOTP

Generate TOTPs with Go. Download the binary, put your secrets into the YAML. Ready to go: Call it with the name of your secret and the token will be copied to your clipboard.

## Installation

Just download the binary and enjoy:

[![GitHub release](https://img.shields.io/github/v/release/jkais/gotp)](https://github.com/jkais/gotp/releases)


## Usage

### List Keys
You can generate a list of all the keys of your secrets with `gotp --`. This is very useful to combine it with things like *wofi* or *rofi*. I run it with `gotp --list | wofi --dmenu --prompt "TOPT" | xargs gotp`.

### Generate token
Run it with `gotp <key>`, where *<key>* is the key from your *secrets.yaml*. If possible, the token will be copied to your clipboard. If not, it will be printed to stdout.
