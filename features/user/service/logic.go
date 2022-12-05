package service

import (
	"be13/project/features/user"
	"errors"

	"github.com/go-playground/validator/v10"
)

// bisnis Logic
type UserService struct {
	userRepository user.RepositoryEntities //data repository dri entities
	validate       *validator.Validate
}

func New(repo user.RepositoryEntities) user.ServiceEntities { //dengan kembalian user.service
	return &UserService{
		userRepository: repo,
		validate:       validator.New(),
	}
}

// Create implements user.ServiceEntities
func (service *UserService) Create(input user.CoreUser) (err error) {

	if validateERR := service.validate.Struct(input); validateERR != nil {
		return validateERR
	}

	_, errCreate := service.userRepository.Create(input)
	if errCreate != nil {
		return errors.New("GAGAL MENAMBAH DATA , QUERY ERROR")
	}

	return nil
}

// DeleteById implements user.ServiceEntities
// func (*UserService) DeleteById(id int) error {
// 	panic("unimplemented")
// }

// // GetAll implements user.ServiceEntities
// func (*UserService) GetAll() (data []user.CoreUser, err error) {
// 	panic("unimplemented")
// }

// // GetById implements user.ServiceEntities
// func (*UserService) GetById(id int) (data user.CoreUser, err error) {
// 	panic("unimplemented")
// }

// // Update implements user.ServiceEntities
// func (*UserService) Update(id int, input user.CoreUser) error {
// 	panic("unimplemented")
// }
