package http

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"html/template"
	"image"
	"image/jpeg"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/labstack/echo"
)

func upload(c echo.Context) error {
	// Read form fields
	name := c.FormValue("name")
	discription := c.FormValue("discription")

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

	return c.HTML(http.StatusOK, fmt.Sprintf("<p>Uploaded successfully %d files with fields name=%s and discription=%s.</p>", len(files), name, discription))
}

var templates = template.Must(template.ParseFiles("templates/index.html"))

func show(c echo.Context) error {
	return c.HTML(http.StatusOK, "OK")
}
func renderTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
	if err := templates.ExecuteTemplate(w, tmpl+".html", data); err != nil {
		log.Fatalln("Unable to execute template.")
	}
}

func ShowHandler(w http.ResponseWriter, r *http.Request) {
	file, err := os.Open("/tmp/test.jpg")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	writeImageWithTemplate(w, "show", &img)
}

func writeImageWithTemplate(w http.ResponseWriter, tmpl string, img *image.Image) {
	buffer := new(bytes.Buffer)
	if err := jpeg.Encode(buffer, *img, nil); err != nil {
		log.Fatalln("Unable to encode image.")
	}
	str := base64.StdEncoding.EncodeToString(buffer.Bytes())
	data := map[string]interface{}{"Title": tmpl, "Image": str}
	renderTemplate(w, tmpl, data)
}
