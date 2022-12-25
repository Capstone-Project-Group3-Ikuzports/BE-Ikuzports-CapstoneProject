package delivery

import (
	"ikuzports/features/chat"
	"ikuzports/middlewares"
	"ikuzports/utils/helper"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type ChatDelivery struct {
	chatService chat.ServiceInterface
}

func New(service chat.ServiceInterface, e *echo.Echo) {
	handler := &ChatDelivery{
		chatService: service,
	}

	e.GET("/chats", handler.GetAll, middlewares.JWTMiddleware())
	e.POST("/chats", handler.Create, middlewares.JWTMiddleware())
	e.DELETE("/chats/:id", handler.Delete, middlewares.JWTMiddleware())

}

func (delivery *ChatDelivery) GetAll(c echo.Context) error {
	results, err := delivery.chatService.GetAll()
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("error read data"))
	}
	dataRespon := fromCoreList(results)
	return c.JSON(http.StatusOK, helper.SuccessWithDataResponse("success read all group chat", dataRespon))
}

func (delivery *ChatDelivery) Create(c echo.Context) error {
	chatInput := InsertRequest{}
	errBind := c.Bind(&chatInput)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Add new chat, semua field harus diisi"+errBind.Error()))
	}
	dataCore := toCore(chatInput)
	err := delivery.chatService.Create(dataCore)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponse("internal server error"+err.Error()))
	}
	return c.JSON(http.StatusCreated, helper.SuccessResponse("Success create chat"))
}

func (delivery *ChatDelivery) Delete(c echo.Context) error {
	idParam := c.Param("id")
	id, errConv := strconv.Atoi(idParam)
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Error. Id must integer."))
	}
	userId := middlewares.ExtractTokenUserId(c)
	if userId < 1 {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Failed load user id from JWT token, please check again."))
	}

	err := delivery.chatService.Delete(id, userId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponse("internal server error"+err.Error()))
	}
	return c.JSON(http.StatusCreated, helper.SuccessResponse("success delete chat"))
}
