package nr

import (
	"fmt"
	"github.com/newrelic/newrelic-client-go/v2/newrelic"
	"log"
	utilenv "test/e2e/util/env"
)

func NewClient() *newrelic.NewRelic {
	apiKey := utilenv.GetNrApiKey()
	apiBaseUrl := utilenv.GetNrApiBaseUrl()
	restApiUrl := fmt.Sprintf("%s/v2", apiBaseUrl)
	nerdGraphUrl := fmt.Sprintf("%s/graphql", apiBaseUrl)
	client, err := newrelic.New(
		newrelic.ConfigPersonalAPIKey(apiKey),
		newrelic.ConfigBaseURL(restApiUrl),
		newrelic.ConfigNerdGraphBaseURL(nerdGraphUrl),
	)
	if err != nil {
		log.Panicf("Couldn't create NewRelic client: %s", err)
	}
	return client
}
