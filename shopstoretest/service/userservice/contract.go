package userservice

import "shopstoretest/entity"

type Service struct {
	Repository Repository
}

type Repository interface {
	Register(user entity.User) (entity.User, error)
	IsPhoneNumberUnique(phoneNumber string) (bool, error)
}
