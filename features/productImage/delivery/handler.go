package delivery

import (
	"errors"
	"ikuzports/features/productImage"
	"ikuzports/middlewares"
	"ikuzports/utils/helper"
	"ikuzports/utils/thirdparty"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type ProductImageDelivery struct {
	productImageService productImage.ServiceInterface
}

func New(service productImage.ServiceInterface, e *echo.Echo) {
	handler := &ProductImageDelivery{
		productImageService: service,
	}
	e.POST("/products_images", handler.Create, middlewares.JWTMiddleware())
	e.GET("/products_images", handler.GetAll)
	e.GET("/products_images/:id", handler.GetByID)
}

func (delivery *ProductImageDelivery) Create(c echo.Context) error {
	productImageInput := ProductImageRequest{}
	errBind := c.Bind(&productImageInput)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Error binding data. "+errBind.Error()))
	}

	dataCore := toCore(productImageInput)

	file, _ := c.FormFile("product_image")
	if file != nil {
		res, err := thirdparty.UploadProfile(c, "product_image")
		if err != nil {
			return errors.New("Failed. Cannot Upload Data.")
		}
		log.Print(res)
		dataCore.URL = res
	} else {
		dataCore.URL = "https://www.hostpapa.com/knowledgebase/wp-content/uploads/2018/04/1-13.png"
	}

	err := delivery.productImageService.Create(dataCore)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponse("Failed insert data "+err.Error()))
	}
	return c.JSON(http.StatusOK, helper.SuccessResponse("Succes create data"))

}

func (delivery *ProductImageDelivery) GetAll(c echo.Context) error {
	result, err := delivery.productImageService.GetAll()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponse("failed get all product images data"))
	}

	return c.JSON(http.StatusOK, helper.SuccessWithDataResponse("Success get all product image data", result))
}

func (delivery *ProductImageDelivery) GetByID(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	result, err := delivery.productImageService.GetByID(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponse("failed get all product images data"))
	}

	return c.JSON(http.StatusOK, helper.SuccessWithDataResponse("Success get all product image data", result))
}
