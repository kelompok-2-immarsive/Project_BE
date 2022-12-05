package factory

import (
	userDelivery "be13/project/features/user/delivery"
	userRepo "be13/project/features/user/repository" //alias
	userService "be13/project/features/user/service"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitFactory(e *echo.Echo, db *gorm.DB) {
	userRepofaktory := userRepo.New(db) //menginiasialisasi func new yang ada di repository
	userServiceFaktory := userService.New(userRepofaktory)
	userDelivery.New(userServiceFaktory, e)
}
