package delivery

import (
	"errors"
	"ikuzports/features/user"
	"ikuzports/middlewares"
	"ikuzports/utils/helper"
	"ikuzports/utils/thirdparty"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

type UserDelivery struct {
	userService user.ServiceInterface
}

func New(service user.ServiceInterface, e *echo.Echo) {
	handler := &UserDelivery{
		userService: service,
	}

	e.GET("/users", handler.GetAll, middlewares.JWTMiddleware())
	e.GET("/users/:id", handler.GetById, middlewares.JWTMiddleware())
	e.POST("/users", handler.Create)
	e.PUT("/users/:id", handler.Update, middlewares.JWTMiddleware(), middlewares.UserOnlySameId)
	e.DELETE("/users/:id", handler.Delete, middlewares.JWTMiddleware(), middlewares.UserOnlySameId)
	e.GET("/users/:id/clubs", handler.GetClubs, middlewares.JWTMiddleware())
	e.GET("/users/:id/products", handler.GetProducts, middlewares.JWTMiddleware())
	e.GET("/users/:id/events", handler.GetEvents, middlewares.JWTMiddleware())
	e.GET("/users/:id/transactions", handler.GetTransactions, middlewares.JWTMiddleware())

	//middlewares.IsAdmin = untuk membatasi akses endpoint hanya admin
	//middlewares.UserOnlySameId = untuk membatasi akses user mengelola data diri sendiri saja

}

func (delivery *UserDelivery) GetAll(c echo.Context) error {
	results, err := delivery.userService.GetAll()
	if err != nil {
		if strings.Contains(err.Error(), "Get data success. No data.") {
			return c.JSON(http.StatusOK, helper.SuccessWithDataResponse(err.Error(), results))
		}
		return c.JSON(http.StatusBadRequest, helper.FailedResponse(err.Error()))
	}

	dataResponse := fromCoreList(results)

	return c.JSON(http.StatusOK, helper.SuccessWithDataResponse("Success read all data", dataResponse))
}

func (delivery *UserDelivery) GetById(c echo.Context) error {
	idParam := c.Param("id")
	id, errConv := strconv.Atoi(idParam)
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Error. Id must integer."))
	}
	results, err := delivery.userService.GetById(id)
	if err != nil {
		if strings.Contains(err.Error(), "Get data success. No data.") {
			return c.JSON(http.StatusOK, helper.SuccessWithDataResponse(err.Error(), results))
		}
		return c.JSON(http.StatusBadRequest, helper.FailedResponse(err.Error()))
	}

	dataResponse := fromCore(results)

	return c.JSON(http.StatusOK, helper.SuccessWithDataResponse("Success read user.", dataResponse))
}

func (delivery *UserDelivery) Create(c echo.Context) error {
	userInput := InsertRequest{}
	errBind := c.Bind(&userInput) // menangkap data yg dikirim dari req body dan disimpan ke variabel
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Error binding data. "+errBind.Error()))
	}

	dataCore := toCore(userInput)
	// upload foto
	file, _ := c.FormFile("user_image")
	if file != nil {
		res, err := thirdparty.UploadProfile(c, "user_image")
		if err != nil {
			return errors.New("Failed. Cannot Upload Data.")
		}
		dataCore.UserImage = res
	} else {
		dataCore.UserImage = "https://www.hostpapa.com/knowledgebase/wp-content/uploads/2018/04/1-13.png"
	}
	err := delivery.userService.Create(dataCore)
	if err != nil {
		if strings.Contains(err.Error(), "Error:Field validation") {
			return c.JSON(http.StatusBadRequest, helper.FailedResponse("Some field cannot empty. Details : "+err.Error()))
		}
		if strings.Contains(err.Error(), "Please pick another email.") {
			return c.JSON(http.StatusBadRequest, helper.FailedResponse("Failed insert data "+err.Error()))
		}
		return c.JSON(http.StatusInternalServerError, helper.FailedResponse("Failed insert data. "+err.Error()))
	}
	return c.JSON(http.StatusCreated, helper.SuccessResponse("Success create data"))
}

func (delivery *UserDelivery) Update(c echo.Context) error {
	idParam := c.Param("id")
	id, errConv := strconv.Atoi(idParam)
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Error. Id must integer."))
	}

	userInput := InsertRequest{}
	errBind := c.Bind(&userInput) // menangkap data yg dikirim dari req body dan disimpan ke variabel
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Error binding data. "+errBind.Error()))
	}

	dataCore := toCore(userInput)
	// upload foto
	file, _ := c.FormFile("user_image")
	if file != nil {
		res, err := thirdparty.UploadProfile(c, "user_image")
		if err != nil {
			return errors.New("Failed. Cannot Upload Data.")
		}
		userInput.UserImage = res
	} else {
		userInput.UserImage = "https://www.hostpapa.com/knowledgebase/wp-content/uploads/2018/04/1-13.png"
	}

	err := delivery.userService.Update(dataCore, id)
	if err != nil {
		if strings.Contains(err.Error(), "Error:Field validation") {
			return c.JSON(http.StatusBadRequest, helper.FailedResponse("Some field cannot Empty. Details : "+err.Error()))
		}
		return c.JSON(http.StatusInternalServerError, helper.FailedResponse("Failed update data. "+err.Error()))
	}

	return c.JSON(http.StatusCreated, helper.SuccessResponse("Success update data."))
}

func (delivery *UserDelivery) Delete(c echo.Context) error {
	idParam := c.Param("id")
	id, errConv := strconv.Atoi(idParam)
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Error. Id must integer."))
	}
	err := delivery.userService.Delete(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse(err.Error()))
	}

	return c.JSON(http.StatusOK, helper.SuccessResponse("Success delete data."))
}

func (delivery *UserDelivery) GetClubs(c echo.Context) error {
	idParam := c.Param("id")
	id, errConv := strconv.Atoi(idParam)
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Error. Id must integer."))
	}

	// validasi data di proses oleh user ybs
	userId := middlewares.ExtractTokenUserId(c)
	if userId < 1 {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Failed load user id from JWT token, please check again."))
	}

	// if userId != id {
	// 	return c.JSON(http.StatusBadRequest, helper.FailedResponse("Failed process data, data must be yours."))
	// }

	// process
	results, err := delivery.userService.GetClubs(id)
	if err != nil {
		if strings.Contains(err.Error(), "Get data success. No data.") {
			return c.JSON(http.StatusOK, helper.SuccessWithDataResponse(err.Error(), results))
		}
		return c.JSON(http.StatusBadRequest, helper.FailedResponse(err.Error()))
	}

	dataResponse := fromClubList(results)

	return c.JSON(http.StatusOK, helper.SuccessWithDataResponse("Success read user clubs", dataResponse))
}

func (delivery *UserDelivery) GetProducts(c echo.Context) error {
	idParam := c.Param("id")
	id, errConv := strconv.Atoi(idParam)
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Error. Id must integer."))
	}

	// validasi data di proses oleh user ybs
	userId := middlewares.ExtractTokenUserId(c)
	if userId < 1 {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Failed load user id from JWT token, please check again."))
	}

	// if userId != id {
	// 	return c.JSON(http.StatusBadRequest, helper.FailedResponse("Failed process data, data must be yours."))
	// }

	// process
	results, err := delivery.userService.GetProducts(id)
	if err != nil {
		if strings.Contains(err.Error(), "Get data success. No data.") {
			return c.JSON(http.StatusOK, helper.SuccessWithDataResponse(err.Error(), results))
		}
		return c.JSON(http.StatusBadRequest, helper.FailedResponse(err.Error()))
	}

	dataResponse := fromProductList(results)

	return c.JSON(http.StatusOK, helper.SuccessWithDataResponse("Success read users products", dataResponse))
}

func (delivery *UserDelivery) GetEvents(c echo.Context) error {
	idParam := c.Param("id")
	id, errConv := strconv.Atoi(idParam)
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Error. Id must integer."))
	}

	// validasi data di proses oleh user ybs
	userId := middlewares.ExtractTokenUserId(c)
	if userId < 1 {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Failed load user id from JWT token, please check again."))
	}

	// if userId != id {
	// 	return c.JSON(http.StatusBadRequest, helper.FailedResponse("Failed process data, data must be yours."))
	// }

	// process
	results, err := delivery.userService.GetEvents(id)
	if err != nil {
		if strings.Contains(err.Error(), "Get data success. No data.") {
			return c.JSON(http.StatusOK, helper.SuccessWithDataResponse(err.Error(), results))
		}
		return c.JSON(http.StatusBadRequest, helper.FailedResponse(err.Error()))
	}

	dataResponse := fromEventList(results)

	return c.JSON(http.StatusOK, helper.SuccessWithDataResponse("Success read user events", dataResponse))
}

func (delivery *UserDelivery) GetTransactions(c echo.Context) error {
	idParam := c.Param("id")
	id, errConv := strconv.Atoi(idParam)
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Error. Id must integer."))
	}

	// validasi data di proses oleh user ybs
	userId := middlewares.ExtractTokenUserId(c)
	if userId < 1 {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Failed load user id from JWT token, please check again."))
	}

	if userId != id {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Failed process data, data must be yours."))
	}

	// process
	results, err := delivery.userService.GetTransactions(id)
	if err != nil {
		if strings.Contains(err.Error(), "Get data success. No data.") {
			return c.JSON(http.StatusOK, helper.SuccessWithDataResponse(err.Error(), results))
		}
		return c.JSON(http.StatusBadRequest, helper.FailedResponse(err.Error()))
	}

	dataResponse := fromTransactionList(results)

	return c.JSON(http.StatusOK, helper.SuccessWithDataResponse("Success read user.", dataResponse))
}
