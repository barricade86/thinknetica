package crawler

// Search bot. Scans web resources
// Interface defines bot's contract
type Interface interface {
	Scan(url string, depth int) ([]Document, error)
	BatchScan(urls []string, depth int, workers int) (<-chan Document, <-chan error)
}

// Document - Web document\
type Document struct {
	ID    int
	URL   string
	Title string
	Body  string
}
