package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"sync"
	"sync/atomic"
	"time"
)

const (
	metricsRoute = "/v1/metrics"
	tracesRoute  = "/v1/traces"
	logsRoute    = "/v1/logs"
)

type validationStore struct {
	transactions uint32
	startTime    time.Time
}

type ValidationPayload struct {
	DurationInMillis int64  `json:"duration"`
	Transactions     uint32 `json:"transactions"`
}

func getEnv(key, fallback string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		value = fallback
	}
	return value
}

func (store *validationStore) handleRequest(resp http.ResponseWriter, req *http.Request) {
	fmt.Printf("Mock OTLP request received: %s\n", req.URL.Path)

	atomic.AddUint32(&store.transactions, 1)
	fmt.Printf("transactions: %d \n", store.transactions)

	resp.WriteHeader(http.StatusOK)
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	store := validationStore{startTime: time.Now()}

	go func(s *validationStore) {
		defer wg.Done()

		otlpHttp := http.NewServeMux()

		otlpHttp.HandleFunc(metricsRoute, s.handleRequest)
		otlpHttp.HandleFunc(tracesRoute, s.handleRequest)
		otlpHttp.HandleFunc(logsRoute, s.handleRequest)

		otlpServer := &http.Server{
			Addr: fmt.Sprintf("%s:%s",
				getEnv("OTLP_HTTP_HOST", "127.0.0.1"),
				getEnv("OTLP_HTTP_PORT", "4318"),
			),
			Handler: otlpHttp,
		}

		log.Println("Starting otlp mock server...")
		err := otlpServer.ListenAndServe()
		if err != nil {
			fmt.Printf("error starting otlp mock server: %s\n", err)
			os.Exit(1)
			return
		}
	}(&store)

	go func(s *validationStore) {
		defer wg.Done()

		validationHttp := http.NewServeMux()

		validationHttp.HandleFunc("/validate", func(resp http.ResponseWriter, req *http.Request) {
			resp.Header().Set("Content-Type", "application/json")

			fmt.Printf("transactions collected: %d \n", s.transactions)

			duration := time.Now().Sub(s.startTime)

			if err := json.NewEncoder(resp).Encode(ValidationPayload{duration.Milliseconds(), store.transactions}); err != nil {
				log.Printf("Unable to write response: %v", err)
			}
		})

		validationHttp.HandleFunc("/", func(writer http.ResponseWriter, req *http.Request) {
			if req.URL.Path != "/" {
				fmt.Printf("Mock OTLP request received: %s\n", req.URL.Path)

				writer.WriteHeader(404)

				log.Println(req.URL.Path)
				return
			} else {
				if _, err := io.WriteString(writer, "success"); err != nil {
					log.Printf("Unable to write response: %v", err)
				}
			}
		})

		validationServer := &http.Server{
			Addr: fmt.Sprintf("%s:%s",
				getEnv("VALIDATION_HTTP_HOST", "127.0.0.1"),
				getEnv("VALIDATION_HTTP_PORT", "8080"),
			),
			Handler: validationHttp,
		}

		log.Println("Starting validation server...")

		err := validationServer.ListenAndServe()
		if err != nil {
			fmt.Printf("error starting validation server: %s\n", err)
			os.Exit(1)
			return
		}

	}(&store)

	wg.Wait()
}
