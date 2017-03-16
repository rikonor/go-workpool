go-workpool
---

#### Usage

```go
// Create a new worker pool with 3 workers
wp := NewWorkerPool(3)

// define a work function
doWork := func() error {
  ...
  // do stuff
  ...
}

// Submit some work for processing
err := wp.SubmitWork(doStuff)
if err != nil {
  log.Fatal(err)
}

// Submit lots of work
for _, wfn := range workFunctions {
  if err := wp.SubmitWork(wfn); err != nil {
    log.Fatal(err)
  }
}

// Wait for the worker pool to finish processing their work
err := wp.Wait()
if err != nil {
  log.Fatal(err)
}
```
