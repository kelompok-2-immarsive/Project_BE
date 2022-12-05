package user

import "time"

type CoreUser struct {
	ID        uint
	FullName  string
	Password  string `validate:"required"`
	Email     string `validate:"required,email"`
	Phone     string `validate:"required"`
	Address   string `validate:"required"`
	Role      string `validate:"required"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type ServiceEntities interface { //sebagai contract yang dibuat di layer service
	// GetAll() (data []CoreUser, err error) //yang returnnya(mengembalikan data core)
	Create(input CoreUser) (err error) // menambahkah data user berdasarkan data usercore
	// Update(id int, input CoreUser) error
	// GetById(id int) (data CoreUser, err error)
	// DeleteById(id int) error
}

type RepositoryEntities interface { // berkaitan database
	// GetAll() (data []CoreUser, err error)
	Create(input CoreUser) (row int, err error)
	// Update(id int, input CoreUser) error
	// GetById(id int) (data CoreUser, err error)
	// DeleteById(id int) error
}
