## TCP and UDP port numbers
Well known ports parsed from https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers

ports.json - JSON array file with all ports parsed from Wikipedia in format:
```go
type Port struct {
    // Start
    // Port number range start and end. Single port is represented as [port, port], i.e. {22, 22}
    // multiple ports are represented as [min, max] included, i.e. {2001, 2009}
    Start uint16 `json:"start"`
    End   uint16 `json:"end"`

    Category    PortCategory `json:"category"`
    Description string       `json:"description"`

    Types map[PortProto]RegistrationStatus `json:"types"`
}
```

## Usage
```shell
# Install
go get github.com/teadove/netports
```

```go
package main

import "github.com/teadove/netports"

func main(){
    fmt.Printf("%+v\n", netports.KnownPorts()[0])
}
```
