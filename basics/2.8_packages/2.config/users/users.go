package users

//User defines a user object
type User struct {
	FullName string
	Username string
	Password string
}

//New creates a new empty user object
func New() User {
	return User{}
}
