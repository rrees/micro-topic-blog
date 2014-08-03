package app

import "fmt"
import "net/http"

func init() {

	http.HandleFunc("/", frontPage)
	
}

func frontPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world")

}