package env

import (
	"fmt"
	"os"
	"strconv"
)

const (
	K8sContextName = "E2E_TEST__K8S_CONTEXT_NAME"
	ImageTag       = "E2E_TEST__IMAGE_TAG"
	NrBackendUrl   = "E2E_TEST__NR_BACKEND_URL"
	NrIngestKey    = "E2E_TEST__NR_INGEST_KEY"
	NrApiKey       = "E2E_TEST__NR_API_KEY"
	NrAccountId    = "E2E_TEST__NR_ACCOUNT_ID"
)

func getEnvVar(envVar string) string {
	value := os.Getenv(envVar)
	if value == "" {
		panic(fmt.Sprintf("%s not set", envVar))
	}
	return value
}

func GetK8sContextName() string {
	return getEnvVar(K8sContextName)
}

func GetImageTag() string {
	return getEnvVar(ImageTag)
}

func GetNrBackendUrl() string {
	return getEnvVar(NrBackendUrl)
}

func GetNrIngestKey() string {
	return getEnvVar(NrIngestKey)
}

func GetNrApiKey() string {
	return getEnvVar(NrApiKey)
}

func GetNrAccountId() int {
	accountIdStr := getEnvVar(NrAccountId)
	accountId, err := strconv.Atoi(accountIdStr)
	if err != nil {
		panic(fmt.Sprintf("Invalid accountId: %s", accountIdStr))
	}
	return accountId
}
