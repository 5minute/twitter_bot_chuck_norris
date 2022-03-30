package main

import (
	"fmt"

	"encoding/json"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"io/ioutil"
	"net/http"
)

type Joke struct {
	Value string `json:"value"`
}

func getQuote() string {
	resp, err := http.Get("https://api.chucknorris.io/jokes/random")
	if err != nil {
		panic(err.Error())
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err.Error())
	}
	var joke Joke
	if err := json.Unmarshal(body, &joke); err != nil {
		panic(err)
	}
	return fmt.Sprintf("Random Chuck joke: %s", joke.Value)
}

func main() {
	consumerKey := "K2zmzBbd0lZhhh7KjoXPjxoq9"
	consumerSecret := "wbmNAtqyo6Qq5zLjPkLCxB2ER95CPPoHwyh6pYiE6tL1BoUc7a"
	accessToken := "1490739894821478404-LjLFN9DfUIbyBChTQ4AX8lmbaNjGzc"
	accessTokenSecret := "TKWe6SCqSGqOBsSYSsCHntp74jRtbXlaMQ1S2ShYipXQe"
	config := oauth1.NewConfig(consumerKey, consumerSecret)
	token := oauth1.NewToken(accessToken, accessTokenSecret)

	httpClient := config.Client(oauth1.NoContext, token)
	client := twitter.NewClient(httpClient)

	user, _, err := client.Accounts.VerifyCredentials(&twitter.AccountVerifyParams{})
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	fmt.Printf("Account: @%s (%s)\n", user.ScreenName, user.Name)
	_, _, err = client.Statuses.Update(getQuote(), nil)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	fmt.Println("Twitted successfully")
}
