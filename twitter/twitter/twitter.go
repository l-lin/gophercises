package twitter

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/l-lin/gophercises/twitter/config"
	"github.com/l-lin/gophercises/twitter/user"
)

const (
	twitterURL  = "https://api.twitter.com/1.1/statuses/retweets/"
	count       = "100"
	twitterUser = "Ligne%d_RATP"
)

var client = &http.Client{Timeout: time.Second * 10}

// Retweet of a tweet
type Retweet struct {
	ID        string `json:"id_str"`
	user.User `json:"user"`
}

// RetweetsResult of the request to find the retweets
type RetweetsResult struct {
	Retweets []Retweet
	Error    error
}

// GetUniqueUsers filter unique users from the retweets
func (r *RetweetsResult) GetUniqueUsers() []user.User {
	users := make([]user.User, len(r.Retweets))
	for i, r := range r.Retweets {
		users[i] = r.User
	}
	return user.UniqueUsers(users)
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
		result <- &RetweetsResult{Retweets: nil, Error: fmt.Errorf("Error status was %d and response body was '%s'", resp.StatusCode, string(body))}
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
