package main

// onyi.wait.oh

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var tmpl *template.Template

type Post struct {
	Title string
	Body  string
	Love  int
}

type AllPostPage struct {
	Title string
	Posts []Post
}

func allPost(w http.ResponseWriter, r *http.Request) {
	all_post := AllPostPage{
		Title: "My Blog",
		Posts: []Post{
			{Title: "Golang Week one", Body: "There is nothing much here to see again", Love: 4},
			{Title: "Golang Week one", Body: "There is nothing much here to see again", Love: 1},
			{Title: "Golang Week one", Body: "There is nothing much here to see again", Love: 3},
		},
	}

	tmpl.Execute(w, all_post)
}

func main() {
	port_address := ":9092"
	fmt.Println("Starting the todo app \n Should be accessible on port " + port_address)
	mux := http.NewServeMux()
	tmpl_dir := "templates"
	tmpl = template.Must(template.ParseFiles(tmpl_dir + "/index.html"))

	fs := http.FileServer(http.Dir("./static"))
	mux.Handle("/static", http.StripPrefix("/static", fs))
	mux.HandleFunc("/", allPost)

	log.Fatal(http.ListenAndServe(":9092", mux))
}
