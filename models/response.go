package models

type Response struct {
	Success bool
	Message string
}

type BookResponse struct {
	Success bool
	Books   []Books
}
