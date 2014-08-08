package app

import "net/http"
import "html/template"

func init() {

	http.HandleFunc("/", frontPage)
	
}


func frontPage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	var frontTemplate = template.Must(template.ParseFiles("templates/base.html", "templates/index.html"))

	templateValues := make(map[string]interface{})

	if err := frontTemplate.Execute(w, templateValues); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}