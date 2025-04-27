package viewx

import (
	"net/http"
	"strings"
	"text/template"

	"github.com/mhhidayat/viewx/templatex"
)

var Title = "Default Title"

func Init(title string) {
	Title = title
	templatex.StartTemplate()
}

func Render(w http.ResponseWriter, viewFile string, data map[string]any) {

	if !strings.HasSuffix(viewFile, ".html") {
		viewFile += ".html"
	}

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
