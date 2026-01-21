package discrete

import "github.com/roidaradal/fn/list"

type (
	Variable = int
	Value    = int
)

// Add the same domain for all variables
func (p *Problem) AddVariableDomains(domain []Value) {
	for _, variable := range p.Variables {
		p.Domain[variable] = domain[:]
	}
}

//////////////// VARIABLE FUNCTIONS ///////////////////

// Create list of variables, mirroring the list of items
func Variables[T any](items []T) []Variable {
	return list.NumRange(0, len(items))
}

// Create list of variables from [0, numItems)
func IndexVariables(numItems int) []Variable {
	return list.NumRange(0, numItems)
}

// Create list of variables from [first, last]
func RangeVariables(first, last int) []Variable {
	return list.NumRange(first, last+1)
}

//////////////// DOMAIN FUNCTIONS  ////////////////////

// Create list of values, mirroring the list of items
func Domain[T any](items []T) []Value {
	return list.NumRange(0, len(items))
}

// Create list of values from [0, numItems)
func IndexDomain(numItems int) []Value {
	return list.NumRange(0, numItems)
}

// Create list of values from [first, last]
func RangeDomain(first, last int) []Value {
	return list.NumRange(first, last+1)
}

// Create boolean domain {1, 0}
func BooleanDomain() []Value {
	return []Value{1, 0}
}

// Create list of values from [-1, 0, numItems)
func PathDomain(numItems int) []Value {
	domain := IndexDomain(numItems)
	domain = append(domain, -1) // not included in path
	return domain
}
