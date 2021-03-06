package http

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/labstack/echo"
)

func upload(c echo.Context) error {
	// Read form fields
	name := c.FormValue("name")
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

	return c.HTML(http.StatusOK, fmt.Sprintf("<p>Uploaded successfully %d files with fields name=%s.</p>", len(files), name))
	//return c.File("templates/show.html")
}

func show(c echo.Context) error {
	files, _ := ioutil.ReadDir("img")
	return c.Render(http.StatusOK, "show.html", files)
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	log.Println(t.templates, name, data)
	return t.templates.ExecuteTemplate(w, name, data)
}

func delete(c echo.Context) error {
	name := c.FormValue("del_name")
	r := os.Remove("img/" + name)
	if r != nil {
		return r
	}
	return c.HTML(http.StatusOK, fmt.Sprintf("<p>Deleted successfully name=%s .</p>", name))

}
