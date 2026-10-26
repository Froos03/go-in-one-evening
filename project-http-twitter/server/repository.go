package server

import "sync"

type tweetsRepository interface {
	AddTweet(t tweet) (int, error)
	Tweets() ([]tweet, error)
}

type TweetsMemoryRepository struct {
	tweets []tweet
	lock   sync.RWMutex
}

func (t *TweetsMemoryRepository) Tweets() ([]tweet, error) {
	t.lock.RLock()
	defer t.lock.RUnlock()
	return t.tweets, nil
}

func (t *TweetsMemoryRepository) AddTweet(tw tweet) (int, error) {
	t.lock.Lock()
	defer t.lock.Unlock()
	t.tweets = append(t.tweets, tw)
	return len(t.tweets), nil
}
