package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
)

//TwitterConfig stores the data from json twitterConfig.json file
type TwitterConfig struct {
	ScreenName        string `json:"screenName"`
	ConsumerKey       string `json:"consumer_key"`
	ConsumerSecret    string `json:"consumer_secret"`
	AccessToken       string `json:"access_token"`
	AccessTokenSecret string `json:"access_token_secret"`
}

var twitterConfig TwitterConfig

func readTwitterConfigTokensAndConnect() (client *twitter.Client) {

	file, e := ioutil.ReadFile("twitterConfig.json")
	if e != nil {
		fmt.Println("error:", e)
	}
	content := string(file)
	json.Unmarshal([]byte(content), &twitterConfig)
	fmt.Println("twitterConfig.json read comlete")

	fmt.Print("connecting to twitter api --> ")
	configu := oauth1.NewConfig(twitterConfig.ConsumerKey, twitterConfig.ConsumerSecret)
	token := oauth1.NewToken(twitterConfig.AccessToken, twitterConfig.AccessTokenSecret)
	httpClient := configu.Client(oauth1.NoContext, token)
	// twitter client
	client = twitter.NewClient(httpClient)

	fmt.Println("connection successful")

	return client
}
