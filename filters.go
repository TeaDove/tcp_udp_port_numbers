package netports

import (
	"iter"
	"slices"
)

type Filter func(Port) bool

func (r Ports) Filter(callbacks ...Filter) iter.Seq[Port] {
	skip := func(port Port) bool {
		for _, callback := range callbacks {
			if !callback(port) {
				return true
			}
		}

		return false
	}

	return func(yield func(Port) bool) {
		for _, port := range r {
			if skip(port) {
				continue
			}

			if !yield(port) {
				return
			}
		}
	}
}

func (r Ports) FilterCollect(callbacks ...Filter) Ports {
	return slices.Collect(r.Filter(callbacks...))
}

func FilterByProto(proto PortProto) Filter {
	return func(port Port) bool {
		status, ok := port.Types[proto]
		if !ok || status == RegistrationNo {
			return false
		}

		return true
	}
}

func FilterByCategory(categories ...PortCategory) Filter {
	return func(port Port) bool {
		return slices.Contains(categories, port.Category)
	}
}

func (r Ports) GroupByNumber() map[uint16]Ports {
	mapping := make(map[uint16]Ports, 1000)

	for _, port := range r {
		portNum := port.Start
		for {
			mapping[portNum] = append(mapping[portNum], port)

			if port.End == portNum {
				break
			}

			portNum++
		}
	}

	return mapping
}
