package main

import (
	"fmt"
	"github.com/briandowns/spinner"
	"time"
)

var (
	apiVersion = "v0.3"
	Version    = "dev"
	BuildDate  = "2021-09-01T00:00:00Z"
	Commit     = "0000"
	ReleaseUrl = "https://api.github.com/repos/cli/cli"
)

// Sum calculates the sum of two integers.
func Sum(a, b int) int {
	return a + b
}

func main() {
	s := spinner.New(spinner.CharSets[9], 100*time.Millisecond)
	s.Suffix = " Waiting for something to happen..."
	s.Start()
	time.Sleep(2 * time.Second)
	s.Stop()

	fmt.Println("Sum of 5 and 7 is:", Sum(5, 7))
	fmt.Println("API version:", apiVersion)
	fmt.Println("Version:", Version)
	fmt.Println("Build date:", BuildDate)
	fmt.Println("Commit:", Commit)
	fmt.Println("Release URL:", ReleaseUrl)
}
