package nr

import (
	"github.com/newrelic/newrelic-client-go/v2/newrelic"
	utilenv "test/e2e/util/env"
)

func NewClient() *newrelic.NewRelic {
	apiKey := utilenv.GetNrApiKey()
	client, err := newrelic.New(newrelic.ConfigPersonalAPIKey(apiKey))
	if err != nil {
		panic(err)
	}
	return client
}
