package storage

import (
	"context"
	"fmt"
	"strings"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

var Domain = "http://lynks.org"

type ShortLinks struct {
	connect *pgxpool.Pool
}

func NewShortLinks(connect *pgxpool.Pool) *ShortLinks {
	return &ShortLinks{connect: connect}
}

func (s *ShortLinks) Add(link string) (string, error) {
	generator := uuid.New()
	uniq := strings.Replace(generator.String(), "-", "", -1)
	_, err := s.connect.Exec(
		context.Background(),
		"INSERT INTO links(`short_link`,`original_link`) VALUES ($1,$2)",
		fmt.Sprintf("%s/%s", Domain, uniq),
		link,
	)

	if err != nil {
		return "", fmt.Errorf("insert error:%s", err)
	}

	return uniq, nil
}

func (s *ShortLinks) FindOriginalByShortLink(shortLink string) (string, error) {
	var originalLink string
	err := s.connect.QueryRow(context.Background(), "select original_link from widgets where short_link=$1 OR original_link=$1", shortLink).Scan(&originalLink)
	if err != nil {
		return "", fmt.Errorf("QueryRow failed: %s\n", err)
	}

	return originalLink, nil
}
