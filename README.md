# bitly-go

Shorten urls with bit.ly in Go and CLI

## Installation

```
go get github.com/yuzuy/bitly-go

// or

go get github.com/yuzuy/bitly-go/cmd/bitly
```

## Usage

### Shortening urls

```
bitly [flags] shorten [url]
```

#### Flags

- t - input your access token. if it not specified, read the token set by "bitly config token"
- v - print the result in detail
- sdomain - input your branded short domain
- sgguid - input your group guid

### Set default token

```
bitly config token [token]
```

*Note: Your token saved in $HOME/.bitly*
