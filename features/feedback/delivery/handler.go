package delivery

import (
	"be13/project/features/feedback"
<<<<<<< HEAD
=======

	"github.com/Labstack/echo"
)

// type FeedbackDelivery struct {
// 	feedbackServices feedback.ServiceInterface
// }
>>>>>>> 23fa500802f7aea9bfdb62ea0d177425fa25d06b

	"github.com/Labstack/echo"
)

type FeedbackDelivery struct {
	feedbackServices feedback.ServiceInterface
}

func NewFeedback(service feedback.ServiceInterface, e *echo.Echo) {
	handler := &FeedbackDelivery{
		feedbackServices: service,
	}

	e.POST("/feedback", handler.AddClass)
	e.PUT("/feedback", handler.UpdateClass)
	e.DELETE("feedback", handler.DeleteClass)

}
