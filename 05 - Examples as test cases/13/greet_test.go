package example

import "fmt"

func ExampleHello() {
	greeting, err := Hello("Jon")
	if err != nil {
		panic(err)
	}
	fmt.Println(greeting)

	// Output:
	// Hello, Jon
}
