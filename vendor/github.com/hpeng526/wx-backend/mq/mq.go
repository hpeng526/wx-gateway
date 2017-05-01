package mq

import "time"

type MessageQueue interface {
	Poll(key string, duration time.Duration) (value string)
	Offer(key string, value string) (err error)
}
