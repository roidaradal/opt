package a

// Config for Assignment problems
type AssignmentCfg struct {
	Tasks      []string
	Workers    []string
	Cost       [][]float64
	Teams      [][]string
	MaxPerTeam int
}
