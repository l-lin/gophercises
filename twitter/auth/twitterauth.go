package auth

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const (
	authURL = "https://api.twitter.com/oauth2/token"
)

var client = &http.Client{Timeout: time.Second * 10}

// TwitterAuthenticator authenticates user to twitter
type TwitterAuthenticator struct {
}

// Authenticate user to twitter
func (a *TwitterAuthenticator) Authenticate(result chan *Result, c *Credentials) {
	v := url.Values{}
	v.Add("grant_type", "client_credentials")
	req, err := http.NewRequest("POST", authURL, strings.NewReader(v.Encode()))
	req.SetBasicAuth(c.UserName, c.Password)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, err := client.Do(req)

	if err != nil {
		result <- &Result{Token: "", Error: err}
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		result <- &Result{Token: "", Error: err}
		return
	}
	if resp.StatusCode != 200 {
		result <- &Result{Token: "", Error: fmt.Errorf("Could not authenticate. Error status was %d and response body was '%s'", resp.StatusCode, string(body))}
		return
	}
	var ar authResponse
	if err := json.Unmarshal(body, &ar); err != nil {
		result <- &Result{Token: "", Error: err}
		return
	}
	result <- &Result{Token: ar.AccessToken, Error: nil}
}

type authResponse struct {
	TokenType   string `json:"token_type"`
	AccessToken string `json:"access_token"`
}
