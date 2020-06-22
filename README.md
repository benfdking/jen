# jen

jen is you go to CLI tool whenever you need to generate jwts for use with a jwks based system.

## Installation

```shell
go install github.com/benfdking/jen
```

## Usage

```
$ jen --help

Jenerate a JWT token

Usage:
  jen [claims json] [flags]
  jen [command]

Available Commands:
  defaults    Return default claims
  help        Help about any command
  jwks        Generate a jwk set with optional private and public key
  list        List the default keys, optionally specifying one only returns one jwks url
  version     Return build information

Flags:
  -c, --claims stringToString   Claims for JWT (default [])
  -d, --defaults                Add default claims (default true)
  -f, --file string             Add claims from JSON file
  -h, --help                    help for jen
  -k, --key string              [abc] jwt key to use (default "a")
  -p, --private string          Use private key to sign jwt

Use "jen [command] --help" for more information about a command.
```

## Credits

- `jen` makes use of the [jwx library](https://github.com/lestrrat-go/jwx)
