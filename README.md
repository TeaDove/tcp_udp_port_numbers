## TCP and UDP port numbers
Well known ports parsed from https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers

ports.json - JSON array file with all ports parsed from Wikipedia in format:
```go
type Port struct {
	// Start
	// Start and end range. Single port is represented as [port, port], i.e. {22, 22}
	// multiple ports are represented as [min, max] included, i.e. {2001, 2009}
	Start uint16 `json:"start"`
	End   uint16 `json:"end"`

	Category    string `json:"category"`
	Description string `json:"description"`

	Types map[PortProto]RegistrationStatus `json:"types"`
}
```

## Usage
```shell
go get github.com/teadove/tcp_udp_port_numbers
```
```go
package main

import "github.com/TeaDove/tcp_udp_port_numbers"

func main(){
    fmt.Printf("%+v\n", tcp_udp_port_numbers.KnownPorts()[0])
}
```
