package delivery

import (
	"ikuzports/features/auth"
	"ikuzports/features/user"
	"ikuzports/utils/helper"
	"ikuzports/utils/thirdparty"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"golang.org/x/oauth2"
)

type AuthHandler struct {
	authService       auth.ServiceInterface
	googleOauthConfig *oauth2.Config
	userService       user.ServiceInterface
}

func New(service auth.ServiceInterface, e *echo.Echo, googleOauthConfig *oauth2.Config, userService user.ServiceInterface) {
	handler := &AuthHandler{
		authService:       service,
		googleOauthConfig: googleOauthConfig,
		userService:       userService,
	}
	e.POST("/auth", handler.Login)
	e.GET("/auth/google", handler.LoginGoogle)
	e.GET("/auth/callback", handler.Callback)
}

var (
	oauthStateString = "random"
)

func (handler *AuthHandler) Login(c echo.Context) error {
	userInput := UserRequest{}
	errBind := c.Bind(&userInput)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("failed to bind data"))
	}

	dataCore := ToCore(userInput)
	result, token, err := handler.authService.Login(dataCore)

	if err != nil {
		if strings.Contains(err.Error(), "Error:Field validation") {
			return c.JSON(http.StatusBadRequest, helper.FailedResponse("Some field cannot Empty. Details : "+err.Error()))
		}
		return c.JSON(http.StatusInternalServerError, helper.FailedResponse("Failed to Login. "+err.Error()))
	}

	return c.JSON(http.StatusOK, helper.SuccessWithDataResponse("Login Success.", FromCore(result, token)))
}

func (handler *AuthHandler) LoginGoogle(c echo.Context) error {
	url := handler.googleOauthConfig.AuthCodeURL(oauthStateString)
	return c.Redirect(http.StatusTemporaryRedirect, url)
}

func (handler *AuthHandler) Callback(c echo.Context) error {
	oauth := handler.googleOauthConfig
	state := c.FormValue("state")
	code := c.FormValue("code")

	content, err := thirdparty.GetUserInfo(oauth, state, code, oauthStateString)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("error read google profile data"))
	}

	result, token, errLog := handler.authService.LoginGoogle(content)
	if errLog != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("error read data"))
	}
	return c.JSON(http.StatusOK, helper.SuccessWithDataResponse("Login Google Success.", FromCore(result, token)))

}
