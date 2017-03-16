package workpool

import "golang.org/x/sync/errgroup"

// WorkerPool is a functional wrapper around a group of workers
// that can be utilized for running work functions
type WorkerPool interface {
	SubmitWork(fn func() error) error
	Wait() error
}

type workerPool struct {
	maxWorkers int

	eg         *errgroup.Group
	inProgress int // number of workers being used at the moment
}

// NewWorkerPool creates a new WorkerPool
func NewWorkerPool(maxWorkers int) WorkerPool {
	return &workerPool{
		maxWorkers: maxWorkers,

		eg:         &errgroup.Group{},
		inProgress: 0,
	}
}

// SubmitWork will submit the given work function to be processed
// if no workers are available it will block
func (p *workerPool) SubmitWork(fn func() error) error {
	if p.inProgress == p.maxWorkers {
		if err := p.eg.Wait(); err != nil {
			return err
		}

		p.eg = &errgroup.Group{}
		p.inProgress = 0
	}

	p.eg.Go(fn)
	p.inProgress++

	return nil
}

// Wait for all submitted work functions to finish processing
func (p *workerPool) Wait() error {
	return p.eg.Wait()
}

// MockWorkerPool is a mockable WorkerPool
type MockWorkerPool struct {
	SubmitWorkFn func(fn func() error) error
	WaitFn       func() error
}

// SubmitWork calls the underlying SubmitWork method
func (p *MockWorkerPool) SubmitWork(fn func() error) error {
	return p.SubmitWorkFn(fn)
}

// Wait calls the underlying Wait method
func (p *MockWorkerPool) Wait() error {
	return p.WaitFn()
}
