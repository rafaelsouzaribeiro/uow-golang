package dto

type AuthorBooks struct {
	AuthorName      string  `json:"author_name"`
	AuthorBio       string  `json:"author_bio"`
	BookName        string  `json:"book_name"`
	BookPrice       float64 `json:"book_price"`
	BookDescription string  `json:"book_description"`
}
