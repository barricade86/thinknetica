package common

type Storage interface {
	Add(link string) (string, error)
	FindOriginalByShortLink(shortLink string) (string, error)
}
