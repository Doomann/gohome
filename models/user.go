package models

type User struct {
	id        int
	firstName string
	lastName  string
}

var (
	users  []*User
	nextId = 1
)

func getUsers() []*User {
	return users
}

func addUser(u User) (User, error) {
	u.id = nextId
	nextId++
	users = append(users, &u)
	return u, nil
}
