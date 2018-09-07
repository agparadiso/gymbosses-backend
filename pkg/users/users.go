package users

type Gym struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type User struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Gym   *Gym   `json:"gym"`
}

type UserSrv interface {
	IsExistingUser(email string) (bool, error)
	SignUp(name, email, gymName string) error
}
