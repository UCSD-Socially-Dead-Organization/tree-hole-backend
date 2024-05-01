package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"os"
	"time"
)

const serverPort = 8001

type Match struct {
	ID        string `json:"id"`
	User1     string `json:"user1"`
	User2     string `json:"user2"`
	CreatedAt string `json:"created_at"`
}

type User struct {
	ID        string `json:"id"`
	Username  string `json:"username"`
	LastLogin string `json:"lastLogin"`
	// CreatedAt string `json:"created_at"`
}

type GetUsersResponse struct {
	Users []User `json:"users"`
}

type GetMatchesResponse struct {
	Matches []Match `json:"matches"`
}

func Request(requestType string, url string, body []byte) (res *http.Response, err error) {
	req, err := http.NewRequest(requestType, url, bytes.NewReader(body))
	if err != nil {
		fmt.Printf("client: could not create request %s\n", err)
		return nil, err
	}

	var token = os.Getenv("TOKEN")

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

	client := http.Client{
		Timeout: 5 * time.Second,
	}

	res, err = client.Do(req)
	if err != nil {
		fmt.Printf("client: error making request %s\n", err)
		return nil, err
	}

	return res, err
}

const charset = "abcdefghijklmnopqrstuvwxyz" +
	"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var seededRand *rand.Rand = rand.New(
	rand.NewSource(time.Now().UnixNano()))

func StringWithCharset(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func GetRandString(length int) string {
	return StringWithCharset(length, charset)
}

func GetActiveUsers() (users []User) {
	url := fmt.Sprintf("http://localhost:%d/v1/users/active", serverPort)
	res, err := Request(http.MethodGet, url, nil)
	if err != nil {
		fmt.Printf("client: error making request %s\n", err)
		os.Exit(1)
	}
	fmt.Printf("client: status code: %d\n", res.StatusCode)

	reqBody, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("client: could not read request body: %s\n", err)
	}
	fmt.Printf("client: request body: %s\n", reqBody)
	parsedReqBody := GetUsersResponse{}
	json.Unmarshal(reqBody, &parsedReqBody)

	return parsedReqBody.Users
}

func GetMatches() (matches []Match) {
	url := fmt.Sprintf("http://localhost:%d/v1/matches", serverPort)
	res, err := Request(http.MethodGet, url, nil)
	if err != nil {
		fmt.Printf("client: error making request %s\n", err)
		os.Exit(1)
	}
	fmt.Printf("client: status code: %d\n", res.StatusCode)

	reqBody, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("client: could not read request body: %s\n", err)
	}
	fmt.Printf("client: request body: %s\n", reqBody)
	parsedReqBody := GetMatchesResponse{}
	json.Unmarshal(reqBody, &parsedReqBody)

	return parsedReqBody.Matches
}

func CreateMatch(user1 string, user2 string) {
	match := Match{
		User1: user1,
		User2: user2,
	}
	jsonBody, err := json.Marshal(match)
	if err != nil {
		fmt.Printf("client: could not marshal json %s\n", err)
		os.Exit(1)
	}

	url := fmt.Sprintf("http://localhost:%d/v1/matches", serverPort)
	res, err := Request(http.MethodPost, url, jsonBody)
	if err != nil {
		fmt.Printf("client: error making request %s\n", err)
		os.Exit(1)
	}
	fmt.Printf("client: status code: %d\n", res.StatusCode)

	reqBody, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("client: could not read request body: %s\n", err)
	}
	fmt.Printf("client: request body: %s\n", reqBody)
	// parsedReqBody := GetMatchesResponse{}
	// json.Unmarshal(reqBody, &parsedReqBody)

	// return parsedReqBody.Matches
}

func CreateUser(username string) {
	user := User{
		Username:  username,
		LastLogin: time.Now().Format(time.RFC3339Nano),
	}

	jsonBody, err := json.Marshal(user)
	if err != nil {
		fmt.Printf("client: could not marshal json %s\n", err)
		os.Exit(1)
	}

	url := fmt.Sprintf("http://localhost:%d/v1/users", serverPort)
	res, err := Request(http.MethodPost, url, jsonBody)
	if err != nil {
		fmt.Printf("client: error making request %s\n", err)
		os.Exit(1)
	}
	fmt.Printf("client: status code: %d\n", res.StatusCode)

	reqBody, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("client: could not read request body: %s\n", err)
	}
	fmt.Printf("client: request body: %s\n", reqBody)
	// parsedReqBody := GetMatchesResponse{}
	// json.Unmarshal(reqBody, &parsedReqBody)

	// return parsedReqBody.Matches
}

func main() {
	// Create 10 fake users
	CreateUser("abchajksdh@test.com")
	for i := 0; i < 10; i++ {
		CreateUser(GetRandString(10) + "@test.com")
	}

	// CreateMatch("abchajksdh@test.com", "defasdasd@test.com")

	// matches := GetMatches()
	// for i, v := range matches {
	// 	fmt.Printf("matches: %d: %s\n", i, v)
	// }

	users := GetActiveUsers()

	rand.Shuffle(len(users), func(i, j int) { users[i], users[j] = users[j], users[i] })
	for i, v := range users {
		fmt.Printf("user: %d: %s\n", i, v)
	}

	for i := 0; i+1 < len(users); i += 2 {
		CreateMatch(users[i].Username, users[i+1].Username)
	}
}
