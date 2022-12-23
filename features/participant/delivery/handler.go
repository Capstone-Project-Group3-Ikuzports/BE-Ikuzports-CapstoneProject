package delivery

import (
	"ikuzports/features/participant"
	"ikuzports/middlewares"
	"ikuzports/utils/helper"
	"net/http"

	"github.com/labstack/echo/v4"
)

type ParticipantDelivery struct {
	participantService participant.ServiceInterface
}

func New(service participant.ServiceInterface, e *echo.Echo) {
	handler := &ParticipantDelivery{
		participantService: service,
	}
	e.POST("/participants", handler.Create, middlewares.JWTMiddleware())
}

func (delivery *ParticipantDelivery) Create(c echo.Context) error {
	participantInput := ParticipantRequest{}
	errBind := c.Bind(&participantInput)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Error binding data. "+errBind.Error()))
	}

	dataCore := toCore(participantInput)
	dataCore.UserID = middlewares.ExtractTokenUserId(c)

	err := delivery.participantService.Create(dataCore)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponse("Failed insert data "+err.Error()))
	}
	return c.JSON(http.StatusOK, helper.SuccessResponse("Succes create data"))
}
