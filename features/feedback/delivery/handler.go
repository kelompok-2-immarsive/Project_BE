package delivery

import (
	"be13/project/features/feedback"

	"github.com/labstack/echo/v4"
)

type FeedbackDelivery struct {
	feedbackServices feedback.ServiceInterface
}

func NewFeedback(service feedback.ServiceInterface, e *echo.Echo) {
	handler := &FeedbackDelivery{
		feedbackServices: service,
	}

	e.POST("/feedback", handler.AddClass)
	// e.PUT("/feedback", handler.UpdateClass)
	// e.DELETE("feedback", handler.DeleteClass)

}
func (delivery *FeedbackDelivery) AddClass(c echo.Context) error {
	// role := middlewares.ExtractTokenUserRole(c)
	// fmt.Println(role)
	// if role != "super admin" {
	// 	return c.JSON(http.StatusUnauthorized, helper.FailedResponse("Failed role is not super admin"))

	// }
	// classs := class.ClassRequest{}
	// errBind := c.Bind(&class)
	// if errBind != nil {
	// 	return c.JSON(http.StatusBadRequest, helper.FailedResponse("error binding"))
	// }

	// result := class.reqToCore()

	// err := delivery.classServices.AddClass(result)
	// if err != nil {
	// 	return c.JSON(http.StatusBadRequest, helper.FailedResponse("error insert into database"))
	// }
	// return c.JSON(http.StatusOK, helper.SuccessResponse("Success Add Class"))

}
