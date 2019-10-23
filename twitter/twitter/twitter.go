package twitter

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/l-lin/gophercises/twitter/config"
)

const (
	twitterURL  = "https://api.twitter.com/1.1/statuses/retweets/"
	count       = "100"
	twitterUser = "Ligne%d_RATP"
)

var client = &http.Client{Timeout: time.Second * 10}

// Retweet of a tweet
type Retweet struct {
	ID   string `json:"id_str"`
	User `json:"user"`
}

// User of a tweet
type User struct {
	ID   string `json:"id_str"`
	Name string `json:"name"`
}

// RetweetsResult of the request to find the retweets
type RetweetsResult struct {
	Retweets []Retweet
	Error    error
}

// GetRetweets for the given tweet id
func GetRetweets(result chan *RetweetsResult, tweetID string) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s%s.json", twitterURL, tweetID), nil)
	if err != nil {
		result <- &RetweetsResult{Retweets: nil, Error: err}
		return
	}
	q := req.URL.Query()
	q.Add("count", count)
	req.URL.RawQuery = q.Encode()
	req.Header.Add("Authorization", "Bearer "+config.GetAPIAuthToken())
	resp, err := client.Do(req)
	if err != nil {
		result <- &RetweetsResult{Retweets: nil, Error: err}
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		result <- &RetweetsResult{Retweets: nil, Error: err}
		return
	}
	if resp.StatusCode != 200 {
		result <- &RetweetsResult{Retweets: nil, Error: fmt.Errorf("Could not authenticate. Error status was %d and response body was '%s'", resp.StatusCode, string(body))}
		return
	}
	var data []Retweet
	err = json.Unmarshal(body, &data)
	if err != nil {
		result <- &RetweetsResult{Retweets: nil, Error: err}
		return
	}
	result <- &RetweetsResult{Retweets: data, Error: nil}
}
