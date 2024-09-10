package gocreatedummyfile

import (
	"crypto/rand"
	"fmt"
	"os"
)

func createLargeFile(filename string, size int64) error {
	// Create or open the file for writing
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("failed to create file: %v", err)
	}
	defer file.Close()

	// Define the size of each chunk to be written
	const chunkSize int64 = 1024 * 1024 // 1 MB
	buffer := make([]byte, chunkSize)

	var written int64
	for written < size {
		// Calculate remaining size to be written
		remaining := size - written
		if remaining < chunkSize {
			// Adjust chunk size for the last write
			buffer = make([]byte, remaining)
		}

		// Write random data to the buffer
		_, err := rand.Read(buffer)
		if err != nil {
			return fmt.Errorf("failed to generate random data: %v", err)
		}

		// Write buffer to the file
		_, err = file.Write(buffer)
		if err != nil {
			return fmt.Errorf("failed to write to file: %v", err)
		}

		written += int64(len(buffer))
	}

	fmt.Printf("File %s created successfully with size %d bytes\n", filename, size)

	return nil
}
