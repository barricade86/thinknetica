package storage

import (
	"context"
	"fmt"
	"strings"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

var Domain = "http://lynks.org"

type PgShortLinks struct {
	connect *pgxpool.Pool
}

func NewPgShortLinks(connect *pgxpool.Pool) *PgShortLinks {
	return &PgShortLinks{connect: connect}
}

func (p *PgShortLinks) Add(link string) (string, error) {
	generator := uuid.New()
	uniq := strings.Replace(generator.String(), "-", "", -1)
	_, err := p.connect.Exec(
		context.Background(),
		"INSERT INTO links(`short_link`,`original_link`) VALUES ($1,$2)",
		fmt.Sprintf("%s/%s", Domain, uniq),
		link,
	)

	if err != nil {
		return "", fmt.Errorf("insert error:%s", err)
	}

	return fmt.Sprintf("%s/%s", Domain, uniq), nil
}

func (p *PgShortLinks) FindOriginalByShortLink(shortLink string) (string, error) {
	var originalLink string
	err := p.connect.QueryRow(context.Background(), "select original_link from widgets where short_link=$1", shortLink).Scan(&originalLink)
	if err != nil {
		return "", fmt.Errorf("QueryRow failed: %s\n", err)
	}

	return originalLink, nil
}
