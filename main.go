package main

import (
	"fmt"
	"ikuzports/config"
	"ikuzports/factory"
	"ikuzports/middlewares"
	"ikuzports/utils/database/mysql"
	"ikuzports/utils/thirdparty"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	cfg := config.GetConfig()
	db := mysql.InitDB(cfg)
	// db := posgresql.InitDB(cfg)
	googleOauthConfig := thirdparty.InitOauth()

	e := echo.New()

	factory.InitFactory(e, db, googleOauthConfig)

	// middleware
	middlewares.LogMiddlewares(e)
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.CORS())
	e.Use(middleware.Logger())

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", cfg.SERVER_PORT)))
}
