package main

import (
	"fmt"
	"github.com/nexmoinc/neru-runtimelib/zip"
	"os"
)

// Sum calculates the sum of two integers.
func Sum(a, b int) int {
	return a + b
}

func main() {
	fmt.Println("Sum of 5 and 7 is:", Sum(5, 7))
	bytes, err := os.ReadFile("testdata/test.zip")
	if err != nil {
		fmt.Println(err)
		return
	}
	if err := uncompressToDir(bytes); err != nil {
		fmt.Println(err)
		return
	}
}

func uncompressToDir(bytes []byte) error {
	if err := zip.UncompressToDir(bytes, "testdata/"); err != nil {
		return err
	}
	return nil
}
