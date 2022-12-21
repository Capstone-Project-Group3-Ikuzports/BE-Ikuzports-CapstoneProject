package middlewares

import (
	"ikuzports/config"
	"ikuzports/utils/helper"
	"net/http"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var key string

func InitJWT(c *config.AppConfig) {
	key = c.JWT_SECRET
}

func JWTMiddleware() echo.MiddlewareFunc {
	return middleware.JWTWithConfig(middleware.JWTConfig{
		SigningMethod: "HS256",
		SigningKey:    []byte(key),
	})
}

func CreateToken(userId int, name string) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["userId"] = userId
	claims["name"] = name
	claims["role"] = "User"
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix() //Token expires after 1 hour
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(key))

}

func ExtractTokenUserId(e echo.Context) int {
	user := e.Get("user").(*jwt.Token)
	if user.Valid {
		claims := user.Claims.(jwt.MapClaims)
		userId := claims["userId"].(float64)
		return int(userId)
	}
	return 0
}

func ExtractTokenUserName(e echo.Context) string {
	user := e.Get("user").(*jwt.Token)
	if user.Valid {
		claims := user.Claims.(jwt.MapClaims)
		role := claims["name"].(string)
		return role
	}
	return ""
}

func IsAdmin(next echo.HandlerFunc) echo.HandlerFunc {
	return func(e echo.Context) error {
		user := e.Get("user").(*jwt.Token)
		if user.Valid {
			claims := user.Claims.(jwt.MapClaims)
			role := claims["role"].(string)

			if role != "Super Admin" {
				return e.JSON(http.StatusUnauthorized, helper.FailedResponse("Error. User unauthorized to access."))
			}
		}
		return next(e)

	}
}

func UserOnlySameId(next echo.HandlerFunc) echo.HandlerFunc {
	return func(e echo.Context) error {
		user := e.Get("user").(*jwt.Token)
		if user.Valid {
			claims := user.Claims.(jwt.MapClaims)

			role := claims["role"].(string)

			// jika role bukan user (super admin) skip fungsi ini
			if role == "User" {
				userIdToken := claims["userId"].(float64)
				idToken := int(userIdToken)

				userIdParam := e.Param("id")
				idParam, errConv := strconv.Atoi(userIdParam)
				if errConv != nil {
					return e.JSON(http.StatusBadRequest, helper.FailedResponse("Error. Id must integer."))
				}

				if idToken != idParam {
					return e.JSON(http.StatusUnauthorized, helper.FailedResponse("Error. User unauthorized to access data other user."))
				}
			}
		}
		return next(e)

	}
}
