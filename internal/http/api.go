package http

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/labstack/echo"
)

func upload(c echo.Context) error {
	// Read form fields
	//name := c.FormValue("name")
	//discription := c.FormValue("discription")

	//------------
	// Read files
	//------------

	// Multipart form
	form, err := c.MultipartForm()
	if err != nil {
		return err
	}
	files := form.File["files"]

	for _, file := range files {
		// Source
		src, err := file.Open()
		if err != nil {
			return err
		}
		defer src.Close()

		// Destination
		dst, err := os.Create("img/" + file.Filename)
		if err != nil {
			return err
		}
		defer dst.Close()

		// Copy
		if _, err = io.Copy(dst, src); err != nil {
			return err
		}
	}

	//return c.HTML(http.StatusOK, fmt.Sprintf("<p>Uploaded successfully %d files with fields name=%s and discription=%s.</p>", len(files), name, discription))
	return c.File("templates/show.html")
}

func show(c echo.Context) error {
	files, _ := ioutil.ReadDir("img")
	for _, f := range files {
		fmt.Println(f.Name())
	}
	return c.HTML(http.StatusOK, "OK")
}
