# Simple comparisons
https://golang.org/ref/spec#Comparison_operators

Don't use `.Equals()` this is not Java. We can compare objects or pointers but keeping in mind the address value and fixing the comparisons using the & or * depending on the case. Be careful with the nil pointers deference errors.
