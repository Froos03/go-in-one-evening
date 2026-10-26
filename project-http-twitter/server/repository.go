package server

import "sync"

type tweetsRepository interface {
	AddTweet(t tweet) (int, error)
	Tweets() ([]tweet, error)
}

type TweetsMemoryRepository struct {
	tweets     []tweet
	tweetsLock sync.RWMutex
}

func (t *TweetsMemoryRepository) Tweets() ([]tweet, error) {
	t.tweetsLock.RLock()
	defer t.tweetsLock.RUnlock()

	return t.tweets, nil
}

func (t *TweetsMemoryRepository) AddTweet(tw tweet) (int, error) {
	t.tweetsLock.Lock()
	defer t.tweetsLock.Unlock()

	t.tweets = append(t.tweets, tw)
	return len(t.tweets), nil
}
