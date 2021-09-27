package main

import (
	"fmt"
	"reflect"
)

func PrintAll(vals []interface{}) {
	for _, val := range vals {
		fmt.Println(reflect.TypeOf(val))
		fmt.Println(val)
	}
}

// This would get:
// "cannot use names (type []string) as type []interface{}"
// func main() {
// 	names := []string{"stanley", "david", "oscar"}
// 	PrintAll(names)
// }

// Correct implementation
func main() {
	names := []string{"stanley", "david", "oscar"}
	vals := make([]interface{}, len(names))
	for i, v := range names {
		vals[i] = v // cast to interface ?
	}
	PrintAll(vals)
}
