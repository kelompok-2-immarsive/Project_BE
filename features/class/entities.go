package class

import "time"

type Core struct {
	ID        uint
	ClassName string
	UserID    uint `validate:"required"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type ServiceInterface interface {
	GetClass() (data []Core, err error)
	AddClass(input Core) error
	GetClassbyId(name string) (Core, error)
	DeleteClass(id int) error
	UpdateClass(id int, input Core) error
}

type RepositoryInterface interface {
	GetClass() (data []Core, err error)
	AddClass(input Core) error
	GetClassbyId(name string) (Core, error)
	DeleteClass(id int) error
	UpdateClass(id int, input Core) error
}
