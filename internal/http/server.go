package http

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type Template struct{}

func Run() {
	e := echo.New()
	// リクエストIDの設定
	e.Use(middleware.RequestID())
	// loggerの設定
	e.Use(middleware.Logger())
	// recoverの設定
	e.Use(middleware.Recover())

	//TODO what's?  "message": "no matching operation was found"
	//// validator
	//spec, err := gen.GetSwagger()
	//if err != nil {
	//	panic(err)
	//}
	//e.Use(middleware2.OapiRequestValidator(spec))

	e.Static("/", "templates")
	e.POST("/upload", upload)

	e.Logger.Fatal(e.Start(":1232"))
}
