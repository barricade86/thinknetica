package storage

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
)

var Domain = "http://lynks.org"

type RedisShortLinks struct {
	connect *redis.Client
}

func NewRedisShortLinks(connect *redis.Client) *RedisShortLinks {
	return &RedisShortLinks{connect: connect}
}

func (r *RedisShortLinks) Add(link string) (string, error) {
	generator := uuid.New()
	uniq := strings.Replace(generator.String(), "-", "", -1)
	duration := 10 * time.Minute
	shortLink := fmt.Sprintf("%s/%s", Domain, uniq)
	err := r.connect.Set(context.Background(), shortLink, link, duration).Err()
	if err != nil {
		return "", fmt.Errorf("Error saving short link:%s", err)
	}

	return fmt.Sprintf("%s/%s", Domain, uniq), nil
}

func (s *RedisShortLinks) FindOriginalByShortLink(shortLink string) (string, error) {
	var originalLink string
	originalLink, err := s.connect.Get(context.Background(), shortLink).Result()
	if err != nil {
		return "", fmt.Errorf("Get data by key failed: %s\n", err)
	}

	return originalLink, nil
}
