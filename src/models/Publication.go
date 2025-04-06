package models

type Publication struct {
	ID         int    `json:"id,omitempty"`
	Title      string `json:"title,omitempty"`
	Content    string `json:"content,omitempty"`
	AuthorID   int    `json:"authorId,omitempty"`
	AuthorNick string `json:"authorNick,omitempty"`
	Likes      int    `json:"likes"`
	CreatedAt  string `json:"createdAt,omitempty"`
}
