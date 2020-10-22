package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

//compile templates on start
var templates = template.Must(template.ParseGlob("templates/*.html"))

//A Page structure
type Page struct {
	Title string
}

//Display the named template
//the "data interface{}" parameter, can be any type of struct. It passes data to the destination html pages
func display(w http.ResponseWriter, tmplName string, data interface{}) {
	templates.ExecuteTemplate(w, tmplName, data)
}

//The handlers will delicate incoming request to display().
func mainHandler(w http.ResponseWriter, r *http.Request) {
	display(w, "main", &Page{Title: "Home"})
}
func videoHandler(w http.ResponseWriter, r *http.Request) {
	display(w, "video", &Page{Title: "Video"})

}


//mux stands for "HTTP request multiplexer" it matches incoming requests against a list of registered routes and calls a handler
//for the route that matches the URL or other conditions
func handleRequest() {
	//declare new router
	myRouter := mux.NewRouter().StrictSlash(true)

	//handlers
	myRouter.HandleFunc("/", mainHandler)
	myRouter.HandleFunc("/video", videoHandler)


	//this line loads up anything in our static directory(css, image, icons)
	myRouter.PathPrefix("/").Handler(http.StripPrefix("/static", http.FileServer(http.Dir("static/"))))

	//setting Port nunber
	log.Fatal(http.ListenAndServe(":8080", myRouter))

}

//main function
func main() {
	fmt.Print("running adelide cup 2021-api...")
	handleRequest()
	fmt.Print("terminating adelaide cup 2021-api...")
}
