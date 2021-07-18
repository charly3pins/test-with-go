Gotchas with closures and parallel tests
- be aware of the values passed into the closure before calling the t.Parallel() that needs to be copied from the "loop" where you're executing otherwise you'll end up having wrong results because the latest value of your table driven test will be the unique used
