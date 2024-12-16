package test

import "testing"

func TagAsSlowTest(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping slow test in short mode: ", t.Name())
	}
}
