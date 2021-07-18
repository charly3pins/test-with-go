Setup and teardown with parallel subtests
- if inside t.Run you put t.Parallel the execution it will queue the execution for that parallel tests later on
- to solve that wrap your tests and subtests that run in parallel with another t.Run used as a "group" and then call the teardown()
