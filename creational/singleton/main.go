// Singleton is a creational design pattern that lets you ensure that a class has only one instance, while providing a global access point to this instance.
package main

import (
	"fmt"
)

func main() {

	for i := 0; i < 30; i++ {
		go getInstance()
	}

	// Scanln is similar to Scan, but stops scanning at a newline and
	// after the final item there must be a newline or EOF.
	fmt.Scanln()
}
