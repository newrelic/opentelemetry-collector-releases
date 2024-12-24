package test

import (
	"github.com/gruntwork-io/terratest/modules/random"
	"strings"
	envutil "test/e2e/util/env"
	"testing"
)

const (
	slowOnly = "slowOnly"
	fastOnly = "fastOnly"
)

func TagAsSlowTest(t *testing.T) {
	if isModeEnabled(fastOnly) {
		t.Skip("Skipping slow test: ", t.Name())
	}
}

func TagAsFastTest(t *testing.T) {
	if isModeEnabled(slowOnly) {
		t.Skip("Skipping fast test: ", t.Name())
	}
}

func isModeEnabled(mode string) bool {
	return strings.Contains(envutil.GetTestMode(), mode)
}

func NewTestId() string {
	return strings.ToLower(random.UniqueId())
}
