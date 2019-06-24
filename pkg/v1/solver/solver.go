package solver

import (
	"errors"
)

const (
	IcanhazipID = "icanhazip"
	IdentID = "ident"
	IfconfigID = "ifconfig"
	)


// Solver public IP solver
type Solver interface {
	Run() (string, error)
}

type SolverManager struct {
	solvers []Solver
}

// NewSolverManager create solver list by solver identity
func NewSolverManager(solvers []string) (*SolverManager,error) {
	m := &SolverManager{
		solvers: make([]Solver, 0),
	}
	for _,item := range solvers  {
		switch item {
		case IcanhazipID:
			m.solvers = append(m.solvers, &IcanhazipCom{})
		case IdentID:
			m.solvers = append(m.solvers, &IdentMe{})
		case IfconfigID:
			m.solvers = append(m.solvers, &IfconfigCo{})
		}

	}
	return m, nil
}


func (m *SolverManager)Run() (ip string, err error) {
	copyed := make([]Solver, len(m.solvers))
	copy(copyed, m.solvers)
	for idx, solver := range copyed {
		ip, err = solver.Run()
		if err == nil {
			return ip, err
		}
		m.solvers = append(m.solvers[:idx], m.solvers[idx+1:]...)
		m.solvers = append(m.solvers, solver)

	}

	return "", errors.New("all solvers can not solve IP")
}