# Lesson 11: When to use Error vs Fatal

After learning the different signaling in tests in the past lesson, the question is when use one or another method. The answer is the more used in Computer Science: 
> It depends. 

That's because there isn't a rule that says when A happens use Error and when not use Fatal. It will depen on what you are testing and what do you consider enough problematic to stop the execution or not.

Checking the test for this lesson, we can see the following points:
- If the NewRequest returns an error, it makes sense to stop there the execution, otherwise the Handler() call will not be possible to perform it
- If the StatusCode is not 200, it doesn't make sense to stop the execution, because maybe is not an error but we've not controlled correctly the status codes
- If the ReadAll returns an error, it makes sense to stop there the execution, otherwise the Unmarshal() call will fail
- If the Unmarshal returns an error, it makes sense to stop there the executoin, otherwise the checks below regarding the structure received will not be possible to perform it
- If the p.Age or p.Name are not what we expected, it doesn't make sense to stop the execution at the first "value check", because maybe the following "value checks" will be wrong too and with the Fatal it would be impossible until you fixed the first, execute test again and maybe the second fail, and so on. Using the Error we can see all at same time in one execution
 
