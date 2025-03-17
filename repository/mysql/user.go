package mysql

import "github.com/pooya-dehghan/entity"

func (d DB) IsPhoneNumberUnique(phoneNumber string) (bool, error) {
	return true, nil
}

func (d DB) RegisterUser(user entity.User) (createdUser entity.User, err error) {
	return user, nil
}
