package example

import "fmt"

// Page will print out a message asking each person who hasn't checked in
// to do so.
func Page(checkIns map[string]bool) {
	for name, checkIn := range checkIns {
		if !checkIn {
			fmt.Printf("Paging %s; please see the front desk to check in.\n", name)
		}
	}
}
