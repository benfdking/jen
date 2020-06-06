# jen

jen is you go to CLI tool whenever you need to generate jwts for use with a jwks based system.

## Installation

```shell
go install github.com/benfdking/jen
```

## Usage

```
$ jen --help

jen is your go to cmd tool for all things jwt

Usage:
  jen [command]

Available Commands:
  defaults    Return default claims
  help        Help about any command
  jwks        Generate a jwk set with optional private and public key
  jwt         Generate a jwt
  list        List the default keys, optionally specifying one only returns one
  version     Return build information

Flags:
  -h, --help   help for jen

Use "jen [command] --help" for more information about a command.
```

## Credits

- `jen` makes heavy use of the [jwx library](https://github.com/lestrrat-go/jwx)
