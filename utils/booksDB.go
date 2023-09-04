package utils

import "BOOK-API/models"

var Books = []models.Book{
	{
		ID:       "1",
		Title:    "To Kill a Mockingbird",
		Author:   "Harper Lee",
		Quantity: 5,
	},
	{
		ID:       "2",
		Title:    "1984",
		Author:   "George Orwell",
		Quantity: 3,
	},
	{
		ID:       "3",
		Title:    "The Great Gatsby",
		Author:   "F. Scott Fitzgerald",
		Quantity: 7,
	},
	{
		ID:       "4",
		Title:    "The Catcher in the Rye",
		Author:   "J.D. Salinger",
		Quantity: 4,
	},
	{
		ID:       "5",
		Title:    "Pride and Prejudice",
		Author:   "Jane Austen",
		Quantity: 6,
	},
}
