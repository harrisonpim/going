package models

type Article struct {
	*Base
	Title    string
	Content  string
	Author   *Author
	AuthorID uint
}
