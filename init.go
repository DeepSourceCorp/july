package main

import (
	"crypto/md5"
	"fmt"
)

func main() {
	fmt.Errorf("Some error string") // raises SCC-ST1005

	a := 10 // skipcq: SCC-SA4006
	a = 20
	fmt.Println(a)
}

// Print message "This is an unexported function"
func unexportedFunction() {
	fmt.Println("This is an unexported function")
}

// Print the message "This is an exported function"
func ExportedFunction() {
	fmt.Println("This is an exported function")
}

func undocumentedFunc() {
	fmt.Println("Undocumented function")
}

func UndocumentedExportedFunc() {
	fmt.Println("Undocumented exported function")
}
