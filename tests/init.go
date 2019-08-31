// The analyzer must go through all packages, including this one
package tests

import "fmt"

// Doc
func unexportedFunction() {
	fmt.Println("unexported fn in tests package")
}

// Doc
func ExportedFunction() {
	fmt.Println("exported fn in tests package")
}

//
func ExportedFunctionWithoutDoc() {
	fmt.Println("exported fn w/o doc")
}
