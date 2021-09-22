package http

import (
	"html/template"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type Template struct {
	templates *template.Template
}

func Run() {
	e := echo.New()
	e.Renderer = &Template{
		templates: template.Must(template.ParseGlob("templates/*.html")),
	}
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

	e.Static("/img", "img")
	e.Static("/", "templates")
	e.POST("/upload", upload)
	e.GET("/show", show)
	e.POST("/delete", delete)
	e.Logger.Fatal(e.Start(":1232"))
}
