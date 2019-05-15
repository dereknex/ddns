package solver

// Solver public IP solver
type Solver interface {
	Run() (string, error)
}
