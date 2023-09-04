package models

// The part in `...` (struct tag) is used to provide metadata for encoding and decoding the struct when working with JSON data.
// The keys are in Pascal case to export them outside this package and make it global.
// If you want a field to be ignored during JSON translation, write it in small case

type Book struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	Quantity int    `json:"quantity"`
}
