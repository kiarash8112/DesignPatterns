package main

import "fmt"

type Tweet struct{}

type TwitterClient struct{}

func (TwitterClient) getRecentTweets(accessToken string) []Tweet {
	fmt.Println("Getting recent tweets")
	return []Tweet{}
}

type oAuth struct{}

func (oAuth) requestToken(appkey, appSecret string) string {
	fmt.Println("Get Request Token")
	return "requestToken"
}

func (oAuth) getAccessToken(requestToken string) string {
	fmt.Println("Get access Token")
	return "accessToken"
}

type twitterAPI struct {
	appKey string
	secret string
}

func (t twitterAPI) getRecentTweets() []Tweet {
	oauth := oAuth{}
	requestToken := oauth.requestToken("appkey", "secretkey")
	accessToken := oauth.getAccessToken(requestToken)

	twitterClient := TwitterClient{}
	tweets := twitterClient.getRecentTweets(accessToken)
	return tweets
}
