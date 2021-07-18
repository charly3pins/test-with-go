// Package example is a package for demonstrating examples in source code.
package example

import "fmt"

type Demo struct{}

func (d Demo) Hello() {}

func Hello(name string) (string, error) {
	return fmt.Sprintf("Hello, %s", name), nil
}
