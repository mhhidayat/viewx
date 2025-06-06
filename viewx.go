package viewx

import (
	"net/http"
	"path/filepath"
	"text/template"

	"github.com/mhhidayat/viewx/helpers"
	"github.com/mhhidayat/viewx/templatex"
)

var Title = "Default Title"

func Init(title string) {
	Title = title
	templatex.StartTemplate()
}

func Render(w http.ResponseWriter, viewFile string, data map[string]any) {

	viewFile = helpers.FormatFileName(viewFile)

	templates := template.Must(template.ParseFiles(
		"views/layout/index.html",
		"views/layout/navbar.html",
		"views/layout/footer.html",
		"views/"+viewFile,
	))

	if data == nil {
		data = map[string]any{}
	}

	if _, exists := data["Title"]; !exists {
		data["Title"] = Title
	}

	err := templates.ExecuteTemplate(w, "index.html", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func RenderHTML(w http.ResponseWriter, viewFile string, data map[string]any) {

	viewFile = helpers.FormatFileName(viewFile)

	if data == nil {
		data = map[string]any{}
	}

	tmpl, err := template.ParseFiles(viewFile)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.ExecuteTemplate(w, filepath.Base(viewFile), data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
