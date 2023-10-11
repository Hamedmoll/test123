package userservice

import (
	"fmt"
	"shopstoretest/entity"
)

type RegisterRequest struct {
	Name        string `json:"name"`
	Password    string `json:"password"`
	PhoneNumber string `json:"phone_number"`
}

type UserInfo struct {
	ID          uint
	Name        string
	PhoneNumber string
}

type RegisterResponse struct {
	user UserInfo
}

func (s Service) Register(req RegisterRequest) (RegisterResponse, error) {

	fmt.Println("user that we want save it", req)
	uniqueness, uErr := s.Repository.IsPhoneNumberUnique(req.PhoneNumber)
	if uErr != nil {
		fmt.Println("here1")

		return RegisterResponse{}, fmt.Errorf("cant check uniqueness %w", uErr)
	}

	if uniqueness == false {
		fmt.Println("here2")

		return RegisterResponse{}, fmt.Errorf("your phone number is not unique")
	}

	newUser := entity.User{
		ID:          0,
		Name:        req.Name,
		Password:    req.Password,
		Credit:      0,
		PhoneNumber: req.PhoneNumber,
	}

	createdUser, cErr := s.Repository.Register(newUser)

	//fmt.Println("\n\n\n\n", newUser)
	if cErr != nil {
		fmt.Println("here3", cErr)

		return RegisterResponse{}, fmt.Errorf("cant save user in database %w", cErr)
	}

	return RegisterResponse{UserInfo{
		ID:          createdUser.ID,
		Name:        createdUser.Name,
		PhoneNumber: createdUser.PhoneNumber,
	}}, nil
}
