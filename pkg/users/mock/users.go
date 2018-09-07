package mock

type UserSrv struct{}

func NewUsersSrv() *UserSrv {
	return &UserSrv{}
}

func (srv *UserSrv) IsExistingUser(email string) (bool, error) {
	return false, nil
}

func (srv *UserSrv) SignUp(name, email string, gymName string) error {
	return nil
}
