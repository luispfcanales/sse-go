package rendertpl

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

//loadTemplates execute and print template
func loadTemplates() {
	tpl = template.Must(template.New("").Funcs(template.FuncMap{
		"iterate":   iterate,
		"urlToCard": urlToCard,
	}).ParseGlob("template/*.html"))

	template.Must(tpl.ParseGlob("template/helpers/*.html"))
	template.Must(tpl.ParseGlob("template/navbar/*.html"))
	//template.Must(tpl.ParseGlob("template/post/*.html"))
	//template.Must(tpl.ParseGlob("template/card/*.html"))
}
func rendertpl(w http.ResponseWriter, name string, data interface{}) {
	err := tpl.ExecuteTemplate(w, name, data)
	if err != nil {
		log.Fatalf("error to execute template : %v", err)
	}
}

//RenderPage render templateLogin
func RenderPage(w http.ResponseWriter, name string, data interface{}) {
	loadTemplates()
	switch name {
	case "login":
		template.Must(tpl.ParseGlob("template/form/*.html"))
	case "reserva", "paquetes-information", "paquetes-pictures", "paquetes-videos":
		template.Must(tpl.ParseGlob("template/card/*.html"))
	case "post":
		template.Must(tpl.ParseGlob("template/post/*.html"))
	}
	rendertpl(w, name, data)
}
func iterate() []int {
	return []int{1, 12, 11, 15, 19}
}
func urlToCard(query string) string {
	return fmt.Sprintf("/paquetes/%s/information", query)
}
