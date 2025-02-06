// pkg/auth/client.go
package auth

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

type AuthClient struct {
	baseURL    string
	httpClient *http.Client
}

type UserResponse struct {
	ID             string `json:"id"`
	OrganizationID string `json:"organization_id"`
	Email          string `json:"email"`
	Role           string `json:"role"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

func NewAuthClient(baseURL string) *AuthClient {
	return &AuthClient{
		baseURL: baseURL,
		httpClient: &http.Client{
			Timeout: time.Second * 10,
		},
	}
}

func (c *AuthClient) ValidateToken(token string) (*UserResponse, error) {
	log.Printf("Validating token: %s", token)

	token = strings.TrimPrefix(token, "Bearer ")

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/validate", c.baseURL), nil)
	if err != nil {
		log.Printf("Error creating request: %v", err)
		return nil, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
	log.Printf("Making request to: %s with Authorization: %s", req.URL.String(), req.Header.Get("Authorization"))

	resp, err := c.httpClient.Do(req)
	if err != nil {
		log.Printf("Error making request: %v", err)
		return nil, err
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	log.Printf("Response status: %d, body: %s", resp.StatusCode, string(body))

	if resp.StatusCode != http.StatusOK {
		var errResp ErrorResponse
		if err := json.Unmarshal(body, &errResp); err != nil {
			return nil, fmt.Errorf("auth service error: status %d", resp.StatusCode)
		}
		return nil, fmt.Errorf("auth service error: %s", errResp.Error)
	}

	var claims UserResponse
	if err := json.Unmarshal(body, &claims); err != nil {
		return nil, err
	}

	return &claims, nil
}
