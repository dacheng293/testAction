package main

import (
	"fmt"
)

const (
	apiVersion = "v0.3"
	version    = "dev"
	buildDate  = "2021-09-01T00:00:00Z"
	commit     = "0000"
	releaseUrl = "https://api.github.com/repos/cli/cli"
)

// Sum calculates the sum of two integers.
func Sum(a, b int) int {
	return a + b
}

func main() {
	fmt.Println("Sum of 5 and 7 is:", Sum(5, 7))
	fmt.Println("API version:", apiVersion)
	fmt.Println("Version:", version)
	fmt.Println("Build date:", buildDate)
	fmt.Println("Commit:", commit)
	fmt.Println("Release URL:", releaseUrl)
}
