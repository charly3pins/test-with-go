# Go's testing/quick package 
https://pkg.go.dev/testing/quick#Check
That package provides you a method that helps you with black box testing. It will call your function you wanna test with random data and use the validator function you provided to find errors or not. That's another way of testing and having helpers and generators for your tests cases but depending on the case could be tricky, specially for the overflowing of the types (the integers in the example provided, as when they overflow they become negative and the test fails...).
