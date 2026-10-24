package server

type tweetsRepository interface {
	AddTweet(t tweet) (int, error)
	Tweets() ([]tweet, error)
}

type TweetsMemoryRepository struct {
	tweets []tweet
}

func (t *TweetsMemoryRepository) Tweets() ([]tweet, error) {
	return t.tweets, nil
}

func (t *TweetsMemoryRepository) AddTweet(tw tweet) (int, error) {
	t.tweets = append(t.tweets, tw)
	return len(t.tweets), nil
}
