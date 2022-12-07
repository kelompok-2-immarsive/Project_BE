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
	e.PUT("/feedback", handler.UpdateClass)
	e.DELETE("feedback", handler.DeleteClass)

}
