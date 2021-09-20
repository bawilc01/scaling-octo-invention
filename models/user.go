package models

type User struct {
	ID        int
	FirstName string
	LastName  string
}

var (
	// users var is slice pointing at User spots in memory
	users  []*User
	nextID = 1
)
