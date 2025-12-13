// Pacakge worker contains common discrete optimization workers
package worker

type Worker interface {
	Run(Solver, Logger)
}
