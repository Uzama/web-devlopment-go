package entities

type Book struct {
	Title      string `json:"title"`
	NoOfPages  int    `json:"no_of_pages"`
	AuthorName string `json:"author_name"`
}
