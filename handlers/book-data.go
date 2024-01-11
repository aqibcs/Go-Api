package handlers

// book struct represents the structure of a book.
type book struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	Quantity int    `json:"quantity"`
}

// books slice contains initial book data.
var books = []book{
	{ID: "1", Title: "Think and Grow Rich", Author: "Napoleon Hill", Quantity: 2},
	{ID: "2", Title: "Brain on Fire", Author: "Susannah Cahalan", Quantity: 5},
	{ID: "3", Title: "Rich Dad Poor Dad", Author: "Robert Kiyosaki", Quantity: 6},
	{ID: "5", Title: "Atomic Habit", Author: "James Clear", Quantity: 9},
}
