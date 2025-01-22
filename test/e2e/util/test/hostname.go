package test

import "fmt"

func NewHostNamePrefix(envName string, deployId string, hostType string) string {
	// only a prefix as helm chart appends hostId
	return fmt.Sprintf("%s-%s-%s", envName, deployId, hostType)
}

const Wildcard = "%"

func NewNrQueryHostNamePattern(envName string, deployId string, hostType string) string {
	hostId := Wildcard
	return fmt.Sprintf("%s-%s-%s-%s", envName, deployId, hostType, hostId)
}
