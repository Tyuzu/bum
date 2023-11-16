package main

import (
    "net/http"
    "log"
    "fmt"
	"html/template"

    "github.com/julienschmidt/httprouter"
)

func ViewPost(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	if r.Method == "GET" {
		fmt.Println(r.URL.Path)
		tmpl.ExecuteTemplate(w, "show.html", ps.ByName("PostId"))
	}
}

var tmpl = template.Must(template.ParseGlob("templates/*.html"))

func main() {
    router := httprouter.New()
    router.GET("/:PostId", ViewPost)
	
	log.Println("Starting Server")
    log.Fatal(http.ListenAndServe(":5678", router))
}
