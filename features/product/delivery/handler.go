package delivery

import (
	"ikuzports/features/product"
	"ikuzports/middlewares"
	"ikuzports/utils/helper"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type ProductDelivery struct {
	productService product.ServiceInterface
}

func New(service product.ServiceInterface, e *echo.Echo) {
	handler := &ProductDelivery{
		productService: service,
	}
	e.GET("/products", handler.GetAll, middlewares.JWTMiddleware())
	e.POST("/products", handler.Create, middlewares.JWTMiddleware())
	e.GET("/products/:id", handler.GetByID, middlewares.JWTMiddleware())
	e.GET("/products/:id/products_images", handler.GetByIDImages, middlewares.JWTMiddleware())
	e.PUT("/products/:id", handler.Update, middlewares.JWTMiddleware())
	e.DELETE("/products/:id", handler.Delete, middlewares.JWTMiddleware())
}

func (delivery *ProductDelivery) GetAll(c echo.Context) error {
	queryItemCategoryID, _ := strconv.Atoi(c.QueryParam("itemcategory_id"))
	queryCity := c.QueryParam("city")
	queryName := c.QueryParam("name")

	helper.LogDebug("\n isi queryItemCategoryID = ", queryItemCategoryID)
	helper.LogDebug("\n isi queryCity = ", queryCity)
	helper.LogDebug("\n isi queryName= ", queryName)

	res, err := delivery.productService.GetAll(queryItemCategoryID, queryCity, queryName)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("error get all product data"))
	}
	dataResp := fromCoreList(res)

	return c.JSON(http.StatusOK, helper.SuccessWithDataResponse("success get all product data", dataResp))
}

func (delivery *ProductDelivery) Create(c echo.Context) error {
	productInput := ProductRequest{}
	errBind := c.Bind(&productInput)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("error binding"))
	}

	dataCore := toCore(productInput)

	dataCore.UserID = middlewares.ExtractTokenUserId(c)

	err := delivery.productService.Create(dataCore)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponse("Failed create data"))
	}

	return c.JSON(http.StatusOK, helper.SuccessResponse("Success create data"))
}

func (delivery *ProductDelivery) GetByID(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	res, err := delivery.productService.GetByID(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("error get product data by id"))
	}
	dataResp := fromCore(res)

	return c.JSON(http.StatusOK, helper.SuccessWithDataResponse("success get product data by id", dataResp))
}

func (delivery *ProductDelivery) GetByIDImages(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	res, err := delivery.productService.GetImages(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("error get product data by id"))
	}
	dataResp := fromCoreListImage(res)

	return c.JSON(http.StatusOK, helper.SuccessWithDataResponse("success get product images by by id product", dataResp))
}

func (delivery *ProductDelivery) Update(c echo.Context) error {
	idParam := c.Param("id")
	id, errConv := strconv.Atoi(idParam)
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Error. Id must integer."))
	}

	productInput := ProductRequest{}
	errBind := c.Bind(&productInput)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("error binding"))
	}

	userId := middlewares.ExtractTokenUserId(c)
	if userId < 1 {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Failed load user id from JWT token, please check again."))
	}

	productData, errGet := delivery.productService.GetByID(id)
	if errGet != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse(errGet.Error()))
	}

	if userId != int(productData.UserID) {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Failed process data, product must be yours."))
	}

	dataCore := toCore(productInput)
	err := delivery.productService.Update(id, dataCore)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponse("Failed update data"))
	}
	return c.JSON(http.StatusOK, helper.SuccessResponse("Success Update data"))

}

func (delivery *ProductDelivery) Delete(c echo.Context) error {
	idParam := c.Param("id")
	id, errConv := strconv.Atoi(idParam)
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Error. Id must integer."))
	}

	userId := middlewares.ExtractTokenUserId(c)
	if userId < 1 {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Failed load user id from JWT token, please check again."))
	}

	productData, errGet := delivery.productService.GetByID(id)
	if errGet != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse(errGet.Error()))
	}

	if userId != int(productData.UserID) {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Failed process data, data must be yours."))
	}

	err := delivery.productService.Delete(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse(err.Error()))
	}

	return c.JSON(http.StatusOK, helper.SuccessResponse("Success delete data."))
}
