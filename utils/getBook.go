package utils

import (
	"BOOK-API/models"
	"errors"
	"fmt"
)

func GetBookById(id string) (*models.Book, error) {
	for i, b := range Books {
		if b.ID == id {
			return &Books[i], nil
		}
	}
	return nil, errors.New("book not found")
}

func GetBookIndexById(id string) (int, error) {
	for i, b := range Books {
		if b.ID == id {
			return i, nil
		}
	}
	return -1, fmt.Errorf("no book found with id: %s", id)
}
