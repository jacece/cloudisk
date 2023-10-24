package model

type User struct {
	Username     string
	Email        string
	LastActiveAt string
	Status       int
}

func UserInsert(user User) bool {
	return true
}
