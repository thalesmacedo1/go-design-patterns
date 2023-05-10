// Singleton is a creational design pattern that lets you ensure that a class has only one instance, while providing a global access point to this instance.

// Here are some common scenarios where you might want to use the singleton pattern:

// 1. Resource management: If you have a limited resource (such as a database connection or a thread pool), you can use the singleton pattern to ensure that only one instance of the resource is created and shared among multiple parts of your program.

// 2. Configuration settings: If you have global configuration settings that need to be accessed from different parts of your program, you can use a singleton to ensure that all parts of the program are using the same instance of the configuration object.

// 3. Logging: If you have a logging class that you want to use throughout your program, you can use a singleton to ensure that all parts of the program are logging to the same file or database.

// 4. Caching: If you have a cache that needs to be shared among different parts of your program, you can use a singleton to ensure that there is only one instance of the cache and that all parts of the program are accessing the same cache.

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
