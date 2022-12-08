package delivery

import (
	"be13/project/features/feedback"
	"be13/project/helper"
	"net/http"
	"strconv"

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
	e.PUT("/feedback", handler.Updatefeedback)
	e.DELETE("/feedback/:id", handler.Deletefeedback)

}

func (delivery *FeedbackDelivery) Addfeedback(c echo.Context) error {

	// roletoken := middlewares.ExtractTokenUserRole(c)
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
func (delivery *FeedbackDelivery) Updatefeedback(c echo.Context) error {

	id, _ := strconv.Atoi(c.Param("id"))

	userInput := FeedbackRequest{}
	errBind := c.Bind(&userInput) // menangkap data yg dikirim dari req body dan disimpan ke variabel
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.PesanGagalHelper("Error binding data "+errBind.Error()))
	}

	dataCore := FeedbackRequestToUserCore(userInput)
	err := delivery.feedbackServices.UpdateFeedback(id, dataCore)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.PesanGagalHelper("Failed update data"+err.Error()))
	}
	return c.JSON(http.StatusCreated, helper.PesanSuksesHelper("success Update data"))
}

func (delivery *FeedbackDelivery) Deletefeedback(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	result := delivery.feedbackServices.DeleteFeedback(id) //memanggil fungsi service yang ada di folder service

	if result != nil {
		return c.JSON(http.StatusBadRequest, helper.PesanGagalHelper("erorr read data"))
	}

	return c.JSON(http.StatusCreated, helper.PesanSuksesHelper("success Hapus data"))
}
