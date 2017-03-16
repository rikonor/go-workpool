package workpool

import "testing"

func TestWorkerPoolNothing(t *testing.T) {
	wp := NewWorkerPool(0)

	if err := wp.Wait(); err != nil {
		t.Fatalf("unexpected failure: %s", err)
	}
}

func TestWorkerPoolSingle(t *testing.T) {
	wp := NewWorkerPool(1)

	if err := wp.SubmitWork(func() error {
		return nil
	}); err != nil {
		t.Fatalf("unexpected failure: %s", err)
	}

	if err := wp.Wait(); err != nil {
		t.Fatalf("unexpected failure: %s", err)
	}
}

func TestWorkerPoolNotFull(t *testing.T) {
	wp := NewWorkerPool(2)

	if err := wp.SubmitWork(func() error {
		return nil
	}); err != nil {
		t.Fatalf("unexpected failure: %s", err)
	}

	if err := wp.Wait(); err != nil {
		t.Fatalf("unexpected failure: %s", err)
	}
}

func TestWorkerPoolMultiple(t *testing.T) {
	wp := NewWorkerPool(3)

	for i := 0; i < 4; i++ {
		if err := wp.SubmitWork(func() error {
			return nil
		}); err != nil {
			t.Fatalf("unexpected failure: %s", err)
		}
	}

	if err := wp.Wait(); err != nil {
		t.Fatalf("unexpected failure: %s", err)
	}
}
