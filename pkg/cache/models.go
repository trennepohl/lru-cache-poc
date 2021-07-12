package cache

import "time"

type Config struct {
	MaxKeyCount int
}

type Value struct {
	Value string
	Key string
	LastUsed time.Time
}