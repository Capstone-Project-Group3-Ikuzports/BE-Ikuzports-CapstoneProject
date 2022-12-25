package delivery

import (
	"ikuzports/features/clubMember"
	"ikuzports/middlewares"
	"ikuzports/utils/helper"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

type ClubMemberDelivery struct {
	clubMemberService clubMember.ServiceInterface
}

func New(service clubMember.ServiceInterface, e *echo.Echo) {
	handler := &ClubMemberDelivery{
		clubMemberService: service,
	}

	e.GET("/members", handler.GetAll, middlewares.JWTMiddleware())
	e.GET("/members/:id", handler.GetById, middlewares.JWTMiddleware())
	e.POST("/members", handler.Create, middlewares.JWTMiddleware())
	e.DELETE("/members/:id", handler.Delete, middlewares.JWTMiddleware())

}

func (delivery *ClubMemberDelivery) GetAll(c echo.Context) error {
	queryStatus := c.QueryParam("status")
	results, err := delivery.clubMemberService.GetAll(queryStatus)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("error read data"))
	}
	dataRespon := fromCoreList(results)
	return c.JSON(http.StatusOK, helper.SuccessWithDataResponse("success read all users", dataRespon))
}

func (delivery *ClubMemberDelivery) GetById(c echo.Context) error {
	idParam := c.Param("id")
	id, errConv := strconv.Atoi(idParam)
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Error. Id must integer."))
	}
	results, err := delivery.clubMemberService.GetById(id)
	if err != nil {
		if strings.Contains(err.Error(), "Get data success. No data.") {
			return c.JSON(http.StatusOK, helper.SuccessWithDataResponse(err.Error(), results))
		}
		return c.JSON(http.StatusBadRequest, helper.FailedResponse(err.Error()))
	}

	dataResponse := fromCore(results)

	return c.JSON(http.StatusOK, helper.SuccessWithDataResponse("Success read user.", dataResponse))
}
func (delivery *ClubMemberDelivery) Create(c echo.Context) error {
	memberInput := MemberRequest{}
	errBind := c.Bind(&memberInput)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Add new member, semua field harus diisi"+errBind.Error()))
	}
	userId := middlewares.ExtractTokenUserId(c)
	if userId < 1 {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Failed load user id from JWT token, please check again."))
	}
	memberInput.UserID = uint(userId)
	dataCore := toCore(memberInput)
	err := delivery.clubMemberService.Create(dataCore)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponse("internal server error"+err.Error()))
	}
	return c.JSON(http.StatusCreated, helper.SuccessResponse("Success create data"))
}

func (delivery *ClubMemberDelivery) Delete(c echo.Context) error {
	idParam := c.Param("id")
	id, errConv := strconv.Atoi(idParam)
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Error. Id must integer."))
	}
	userId := middlewares.ExtractTokenUserId(c)
	if userId < 1 {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Failed load user id from JWT token, please check again."))
	}

	err := delivery.clubMemberService.Delete(id, userId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponse("internal server error"+err.Error()))
	}
	return c.JSON(http.StatusCreated, helper.SuccessResponse("success delete member"))
}
