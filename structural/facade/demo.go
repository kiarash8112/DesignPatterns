package facade

import "fmt"

func main() {
	oauth := oAuth{}
	requestToken := oauth.requestToken("appkey", "secretkey")
	accessToken := oauth.getAccessToken(requestToken)

	twitterClient := TwitterClient{}
	tweets := twitterClient.getRecentTweets(accessToken)
	fmt.Printf("tweets: %v\n", tweets)
}
