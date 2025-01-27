package test

import "fmt"

func NewHostNamePrefix(envName string, deployId string, hostType string) string {
	// TODO: incorporate distro into hostname when generalizing nightly to support multiple distro
	// only a prefix as helm chart appends hostId
	return fmt.Sprintf("%s-%s-%s", envName, deployId, hostType)
}

const Wildcard = "%"

func NewNrQueryHostNamePattern(envName string, deployId string, hostType string) string {
	// TODO: incorporate distro into hostname when generalizing nightly to support multiple distro
	hostId := Wildcard
	return fmt.Sprintf("%s-%s-%s-%s", envName, deployId, hostType, hostId)
}
