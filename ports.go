package main

import (
	_ "embed"
	"encoding/json"
)

//go:embed ports.json
var portsJson []byte

type PortRange struct {
	Range [2]uint16 `json:"range"`
	Ports []Port    `json:"ports"`
}

type Port struct {
	Category    string `json:"category"`
	Description string `json:"description"`

	Types map[PortProto]string `json:"types"`
}

type Ports []PortRange

var KnownPorts Ports

func init() {
	err := json.Unmarshal(portsJson, &KnownPorts)
	if err != nil {
		panic(err)
	}
}

type PortProto string

const (
	TCP  PortProto = "tcp"
	UDP  PortProto = "udp"
	SCTP PortProto = "sctp"
	DCCP PortProto = "dccp"
)

func (r Ports) Mapping(proto PortProto) map[uint16]Port {
	mapping := make(map[uint16]Port, 1000)
	for _, portRange := range r {
		for _, port := range portRange.Ports {
			_, ok := port.Types[proto]
			if !ok {
				continue
			}

			for i := portRange.Range[0]; i <= portRange.Range[1]; i++ {
				mapping[i] = port
			}
		}
	}

	return mapping
}
