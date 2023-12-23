package main

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

// TestSum tests the Sum function.
func TestSum(t *testing.T) {
	total := Sum(5, 5)
	if total != 10 {
		t.Errorf("Sum was incorrect, got: %d, want: %d.", total, 10)
	}
}

func Test_uncompressToDir(t *testing.T) {
	bytes, err := os.ReadFile("testdata/test.zip")
	if err != nil {
		fmt.Println(err)
		return
	}
	err = uncompressToDir(bytes)
	require.NoError(t, err)
	os.Remove("testdata/vcr.yaml")
}
