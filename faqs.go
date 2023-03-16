package main

// onyi.wait.oh

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"regexp"
)

var tmpl *template.Template

type Faq struct {
	ID       int
	Question string
	Answer   string
}

type AllPostPage struct {
	Title string
	Faqs  []Faq
}

var all_posts = []Faq{
	{
		ID:       1,
		Question: "Question One",
		Answer:   "eligendi explicabo ducimus cum inventore eos. blanditiis ullam distinctio qui",
	},
	{
		ID:       2,
		Question: "Question Two",
		Answer:   "eligendi explicabo ducimus cum inventore eos. delectus repellendus perferendis tempora labore odio",
	},
	{
		ID:       3,
		Question: "Question Three",
		Answer:   "eligendi explicabo ducimus cum inventore eos.",
	},
	{
		ID:       4,
		Question: "Question Four",
		Answer:   "eligendi explicabo ducimus cum inventore eos. ea temporibus praesentium",
	},
}

var tmpl_dir = "templates"

func allFaqs(w http.ResponseWriter, r *http.Request) {
	all_post := AllPostPage{
		Title: "Faqs Page",
		Faqs:  all_posts,
	}
	tmpl = template.Must(template.ParseFiles(tmpl_dir + "/index.html"))

	tmpl.Execute(w, all_post)
}

func faqAdmin(w http.ResponseWriter, r *http.Request) {
	tmpl = template.Must(template.ParseFiles(tmpl_dir + "/admin.html"))
	faqs := AllPostPage{
		Title: "Admin Page",
		Faqs:  all_posts,
	}
	tmpl.Execute(w, faqs)
}

func faqEdit(w http.ResponseWriter, r *http.Request) {
	validPath := regexp.MustCompile("^/faqs/(edit|save|view)/([a-zA-Z0-9]+)$")

	m := validPath.FindStringSubmatch(r.URL.Path)
	if m == nil {
		http.NotFound(w, r)
		return
	}
	// return m[2], nil //The title is the second expression
	fmt.Println(m)

	tmpl = template.Must(template.ParseFiles(tmpl_dir + "/edit.html"))

	tmpl.Execute(w, "Joy")
}

func faqCreate(w http.ResponseWriter, r *http.Request) {
	tmpl_dir := "templates"
	tmpl = template.Must(template.ParseFiles(tmpl_dir + "/create.html"))
}

func handlePageNotFound(w http.ResponseWriter, r *http.Request) {
	tmpl_dir := "templates"
	tmpl = template.Must(template.ParseFiles(tmpl_dir + "/404.html"))

	tmpl.Execute(w, nil)
}

func main() {
	port_address := ":9092"
	fmt.Println("Starting the todo app \n Should be accessible on port " + port_address)
	mux := http.NewServeMux()

	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	mux.HandleFunc("/faqs/all", allFaqs)
	mux.HandleFunc("/faqs/admin", faqAdmin)
	mux.HandleFunc("/faqs/create", faqCreate)
	mux.HandleFunc("/faqs/edit/", faqEdit)
	// mux.HandleFunc("/", handlePageNotFound)

	log.Fatal(http.ListenAndServe(":9092", mux))
}
