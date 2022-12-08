package delivery

import (
	"be13/project/features/feedback"
	"be13/project/helper"
	"net/http"

	"github.com/labstack/echo/v4"
)

type FeedbackDelivery struct {
	feedbackServices feedback.ServiceInterface
}

func NewFeedback(service feedback.ServiceInterface, e *echo.Echo) {
	handler := &FeedbackDelivery{
		feedbackServices: service,
	}

	e.POST("/feedback", handler.Addfeedback)
	// e.PUT("/feedback", handler.Updatefeeback)
	// e.DELETE("feedback", handler.Deletefeedback)

}

func (delivery *FeedbackDelivery) Addfeedback(c echo.Context) error {

	// roletoken := middlewares.ExtractTokenUserRole(c)
	// log.Println("Role Token", roletoken)
	// if roletoken != "admin" {
	// 	return c.JSON(http.StatusUnauthorized, helper.PesanGagalHelper("tidak bisa diakses khusus admin!!!"))
	// }

	InputFeedback := FeedbackRequest{}
	errbind := c.Bind(&InputFeedback)

	if errbind != nil {
		return c.JSON(http.StatusBadRequest, helper.PesanGagalHelper("erorr read data"+errbind.Error()))
	}
	dataCore := FeedbackRequestToUserCore(InputFeedback) //data mapping yang diminta create
	errResultCore := delivery.feedbackServices.AddFeedback(dataCore)
	if errResultCore != nil {
		return c.JSON(http.StatusBadRequest, helper.PesanGagalHelper("erorr read data"+errResultCore.Error()))
	}
	return c.JSON(http.StatusCreated, helper.PesanSuksesHelper("berhasil create user"))
}
