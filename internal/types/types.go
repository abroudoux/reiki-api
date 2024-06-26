package types

type Session struct {
	Id string `json:"id"`
	FirstName string `json:"firstName"`
	LastName string `json:"lastName"`
	Email string `json:"email"`
	Date string `json:"date"`
	Status string `json:"status"`
}

type Message struct {
	Id string `json:"id"`
	Email string `json:"email"`
	Object string `json:"object"`
	Message string `json:"message"`
	Date string `json:"date"`
}