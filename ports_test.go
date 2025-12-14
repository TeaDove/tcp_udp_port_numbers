package main

import (
	"fmt"
	"testing"
)

func TestPortRange(t *testing.T) {
	ssh, ok := KnownPorts.Mapping(TCP)[22]
	if !ok {
		t.Error("ssh port not found")
	}

	fmt.Printf("%+v\n", ssh)
}
