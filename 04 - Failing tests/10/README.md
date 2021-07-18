# Lesson 10: Ways to signal test failure

In order to signal when a test fails the `testing` package provides a set of methods that can be useful for that.

The basic ones are the `t.Log` and `t.Logf`. These will be printed only if a test fails. If you want to print them anyway you will need to pass the flag `- v` when executing the `go test` command. The `fmt.println` is not recommended because it will not provide any context info, like the file or the line, it will just print a string, so better not use it. For debug purposes the two methods commented above are more than enough.

For a mark a test as a failed the `testing.T` type provides 2 methods for that. `t.Fail` will mark the test as failed but keep running. `t.FailNow` will fail the execution of the test and stop the execution.

With those methods in mind, ones for logging and others for failing, the type provides the methods that are most used that are basically a combination of the ones checked before. It has the `t.Error` and `t.Errorf`; these are a combo of `t.Log`/`t.Logf` plus `t.Fail`. And then the other methods that provides are the `t.Fatal` and `t.Fatalf`; these are a combo of `t.Log`/`t.Logf` is `t.FailNow()`. 
