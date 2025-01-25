package test

import (
	"strings"
	envutil "test/e2e/util/env"
)

const (
	Wildcard                 = "%"
	hostNameSegmentSeparator = "-"
)

func NewHostNamePrefix(envName string, deployId string, hostType string) string {
	distro := getNormalizedDistro()
	// only a prefix as helm chart appends hostId
	return strings.Join([]string{envName, deployId, distro, hostType}, hostNameSegmentSeparator)
}

func NewNrQueryHostNamePattern(envName string, deployId string, hostType string) string {
	distro := getNormalizedDistro()
	hostId := Wildcard
	return strings.Join([]string{envName, deployId, distro, hostType, hostId}, hostNameSegmentSeparator)
}

func getNormalizedDistro() string {
	// solely to improve readability - no technical necessity
	return strings.Replace(envutil.GetDistro(), hostNameSegmentSeparator, "_", -1)
}
