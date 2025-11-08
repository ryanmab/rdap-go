[![Test](https://github.com/ryanmab/rdap-go/actions/workflows/test.yml/badge.svg)](https://github.com/ryanmab/rdap-go/actions/workflows/test.yml)
[![Coverage](https://api.coveragerobot.com/v1/graph/github/ryanmab/rdap-go/badge.svg?token=96a67a715099de230799f42a24282371d046eeff17ae6995f0)](https://coveragerobot.com)
[![Go Report Card](https://goreportcard.com/badge/github.com/ryanmab/rdap-go)](https://goreportcard.com/report/github.com/ryanmab/rdap-go)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/ryanmab/rdap-go)](https://pkg.go.dev/github.com/ryanmab/rdap-go)

# rdap-go

A fast Go client for performing lookups of DNS records, IPv4 addresses and IPv6 addresses using the Registration Data Access Protocol.

## Usage

```
go get github.com/ryanmab/rdap-go@v0.1.0
```

### Domain Lookups

```go

package main

import (
    "log"
	"github.com/ryanmab/rdap-go/pkg/client"
)

func main() {
	client := client.New()

	response, err := client.LookupDomain("ryanmaber.co.uk")

	if err != nil {
		log.Panic(err)
	}

	log.Printf("Status: %s", response.Status)
}

```

### IPv4 Lookups

```go

package main

import (
	"github.com/ryanmab/rdap-go/pkg/client"
	"log"
)

func main() {
	client := client.New()

	response, err := client.LookupIPv4("8.8.8.8")

	if err != nil {
		log.Panic(err)
	}

	log.Printf("Name: %s", response.Name) // GOGL
}

```

### IPv6 Lookups

```go

package main

import (
	"github.com/ryanmab/rdap-go/pkg/client"
	"log"
)

func main() {
	client := client.New()

	response, err := client.LookupIPv6("2001:4860:4860::8888")

	if err != nil {
		log.Panic(err)
	}

	log.Printf("Name: %s", response.Name) // GOOGLE-IPV6
}

```

## Contributing

Contributions are welcome, and encouraged - simply fork the repository, and make a pull request!
