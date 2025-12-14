package netports

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPortRange(t *testing.T) {
	t.Parallel()

	ssh, ok := KnownPorts.GroupByNumber()[22]
	assert.True(t, ok)
	assert.Len(t, ssh, 1)

	assert.Equal(
		t,
		"Secure Shell (SSH), secure logins, file transfers (scp, sftp) and port forwarding",
		ssh[0].Description,
	)
}

func TestFilters(t *testing.T) {
	t.Parallel()

	grouped := KnownPorts.FilterCollect(FilterByProto(TCP), FilterByCategory(CategoryWellKnown, CategoryRegistered)).
		GroupByNumber()

	_, ok := grouped[65530]
	assert.False(t, ok)
}

func ExamplePorts() {
	fmt.Printf("%d %d",
		len(KnownPorts.FilterCollect(
			FilterByProto(TCP),
			FilterByCategory(CategoryWellKnown, CategoryRegistered),
		)),
		len(KnownPorts.FilterCollect(
			FilterByProto(UDP),
			FilterByCategory(CategoryWellKnown, CategoryRegistered),
		)),
	)
	// Output: 2853 2448
}
