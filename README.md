# go-create-dummy-file

A simple function to create dummy file, useful for testing.

Usage:

```go
package main

import "github.com/thomas-tacquet/go-create-dummy-file"

func main() {
	// Example: Create a file of 5GB size
	err := gocreatedummyfile.CreateLargeFile("large_test_file.bin", 5*1024*1024*1024) // 5 GB
	if err != nil {
		fmt.Println("Error:", err)
	}
}
```

