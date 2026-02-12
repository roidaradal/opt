package discrete

import "github.com/roidaradal/fn/list"

type (
	Variable = int
	Value    = int
)

// AddVariableDomains adds the same domain for all variables
func (p *Problem) AddVariableDomains(domain []Value) {
	for _, variable := range p.Variables {
		p.Domain[variable] = domain[:]
	}
	p.uniformDomain = domain
}

//////////////// VARIABLE FUNCTIONS ///////////////////

// Variables creates a list of variables, mirroring the list of items
func Variables[T any](items []T) []Variable {
	return list.NumRange(0, len(items))
}

// IndexVariables creates a list of variables from [0, numItems)
//func IndexVariables(numItems int) []Variable {
//	return list.NumRange(0, numItems)
//}

// RangeVariables creates a list of variables from [first, last]
func RangeVariables(first, last int) []Variable {
	return list.NumRange(first, last+1)
}

//////////////// DOMAIN FUNCTIONS  ////////////////////

// Domain creates a list of values, mirroring the list of items
func Domain[T any](items []T) []Value {
	return list.NumRange(0, len(items))
}

// IndexDomain creates a list of values from [0, numItems)
func IndexDomain(numItems int) []Value {
	return list.NumRange(0, numItems)
}

// RangeDomain creates a list of values from [first, last]
func RangeDomain(first, last int) []Value {
	return list.NumRange(first, last+1)
}

// BooleanDomain creates a boolean domain {1, 0}
func BooleanDomain() []Value {
	return []Value{1, 0}
}

// PathDomain creates a list of values from [-1, 0, numItems)
func PathDomain(numItems int) []Value {
	domain := IndexDomain(numItems)
	domain = append(domain, -1) // not included in path
	return domain
}
