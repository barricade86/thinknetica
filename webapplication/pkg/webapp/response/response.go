package response

type IndexData struct {
	Token        string `json:"token"`
	PositionList string `json:"positions_list"`
}

type DocData struct {
	Title string `json:"title"`
	Body  string `json:"body"`
	URL   string `json:"url"`
}
