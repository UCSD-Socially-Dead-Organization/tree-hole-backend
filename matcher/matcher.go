package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

const serverPort = 8001
const token = "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCIsImtpZCI6InFMbF8zUFdwYmE5OVA1VllLaG93UiJ9.eyJjcmVkcy1lbWFpbCI6InRlc3R1c2VyQGdtYWlsLmNvbSIsImNyZWRzLW5pY2tuYW1lIjoidGVzdHVzZXIiLCJjcmVkcy1uYW1lIjoidGVzdHVzZXJAZ21haWwuY29tIiwiY3JlZHMtcGljdHVyZSI6Imh0dHBzOi8vcy5ncmF2YXRhci5jb20vYXZhdGFyL2ExOGJmNzg2ZWZiNzZhM2Q1NmVlNjlhM2IzNDM5NTJhP3M9NDgwJnI9cGcmZD1odHRwcyUzQSUyRiUyRmNkbi5hdXRoMC5jb20lMkZhdmF0YXJzJTJGdGUucG5nIiwiaXNzIjoiaHR0cHM6Ly9kZXYtMTJ3eHhydWJ5cjRtN2tkZy51cy5hdXRoMC5jb20vIiwic3ViIjoiYXV0aDB8NjYyYzgzYTBjYjEwODJiOGVhN2YwZWZmIiwiYXVkIjoiaHR0cHM6Ly90cmVlLWhvbGUtYmFja2VuZCIsImlhdCI6MTcxNDI2OTI5MiwiZXhwIjoxNzE0MzU1NjkyLCJndHkiOiJwYXNzd29yZCIsImF6cCI6IkxwRGx6VmhRd2MxQlZBc3hmMmNEUVU5MEtvQkxxd1JIIn0.O4p_R9WE2qz2q6Nd5AEhxORIXUM_nO_PFu1-XkhKKzIWKqGtDNaviPcMSHsfWb3YzT_6Z2OZH8bziwRznRTGTStwIAVfkk3KjMQaloxAvHvAU1h8tsH_8z7S9LgMLcsg0-FzHhT7MlR1t2j_PQ12sNUJzt-WttpWsT8PsXl7ksb5JQzGkEfGPgm-IsyDoJD3NPHsNwSZsnT5nSrzV-xiYK-R9oai0CCDQXPa9AiurLmpptzgTwtZu4fvRN_PtuisWXQt5YcKSg1NZsR0-TqMbMUFwAeZv_UoWj0Ph6SlF7sqDgpc8_MBpNYsQmV2Tf86kDML54cTnCPbAA7x9Pav_Q"

func getMatches() {
	requestURL := fmt.Sprintf("http://localhost:%d/v1/match", serverPort)
	req, err := http.NewRequest(http.MethodGet, requestURL, nil)
	if err != nil {
		fmt.Printf("client: could not create request %s\n", err)
		os.Exit(1)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

	client := http.Client{
		Timeout: 5 * time.Second,
	}

	res, err := client.Do(req)
	if err != nil {
		fmt.Printf("client: error making request %s\n", err)
		os.Exit(1)
	}

	fmt.Printf("client: got response!\n")
	fmt.Printf("client: status code: %d\n", res.StatusCode)
}

func createMatch() {

	jsonBody := []byte(`{"user1": "abc@test.com", "user2": "def@test.com"}`)
	// jsonBody := []byte(`{"user2": "def@test.com", "id": "abcdef"}`)
	// jsonBody := []byte(`{}`)
	bodyReader := bytes.NewReader(jsonBody)

	requestURL := fmt.Sprintf("http://localhost:%d/v1/match", serverPort)
	req, err := http.NewRequest(http.MethodPost, requestURL, bodyReader)
	if err != nil {
		fmt.Printf("client: could not create request %s\n", err)
		os.Exit(1)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

	client := http.Client{
		Timeout: 5 * time.Second,
	}

	res, err := client.Do(req)
	if err != nil {
		fmt.Printf("client: error making request %s\n", err)
		os.Exit(1)
	}

	fmt.Printf("client: got response!\n")
	fmt.Printf("client: status code: %d\n", res.StatusCode)

	reqBody, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("client: could not read request body: %s\n", err)
	}
	fmt.Printf("client: request body: %s\n", reqBody)
}

func main() {
	createMatch()
	// getMatches()
}
