package views

import (
	"bytes"
	"html/template"
	"io"
	"net/http"
	"path/filepath"
)

var (
	LayoutDir    string = "views/layouts/"
	TemplatePath string = "views/"
	TemplateExt  string = ".gohtml"
)

func NewView(layout string, files ...string) *View {
	prependTemplatePath(files)
	appendTemplateExt(files)
	files = append(files, layoutFiles()...)

	t, err := template.ParseFiles(files...)
	if err != nil {
		panic(err)
	}

	return &View{
		Template: t,
		Layout:   layout,
	}
}

type View struct {
	Template *template.Template
	Layout   string
}

func (v *View) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	v.Render(w, nil)
}

func (v *View) Render(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "text/html")

	switch data.(type) {
	case Data:
	default:
		data = Data{
			Yield: data,
		}
	}

	var buf bytes.Buffer

	if err := v.Template.ExecuteTemplate(&buf, v.Layout, data); err != nil {
		http.Error(w, "Something went wrong. If the problem spersists, please email", http.StatusInternalServerError)
		return
	}
	io.Copy(w, &buf)
}

func layoutFiles() []string {
	files, err := filepath.Glob(LayoutDir + "*" + TemplateExt)
	if err != nil {
		panic(err)
	}
	return files
}

func prependTemplatePath(files []string) {
	for index, file := range files {
		files[index] = TemplatePath + file
	}
}

func appendTemplateExt(files []string) {
	for index, file := range files {
		files[index] = file + TemplateExt
	}
}
