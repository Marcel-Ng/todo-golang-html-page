package main

// onyi.wait.oh

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"regexp"
	"strconv"
)

var tmpl *template.Template

type Faq struct {
	ID       int
	Question string
	Answer   string
}

// This is the best data struct to use for the Faqs
type FaqData struct {
	Faqs []Faq `json:"faqs"`
}

type EditPage struct {
	Title string
	Faq   Faq
}

type AllPostPage struct {
	Title string
	Faqs  []Faq
}

var tmpl_dir = "templates"

func getFagsJSON() (*FaqData, error) {
	var faq FaqData
	body, err := os.ReadFile("./data/faqs.json")
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(body, &faq); err != nil {
		return nil, err
	}
	return &faq, nil
}

func allFaqs(w http.ResponseWriter, r *http.Request) {

	v, e := getFagsJSON()
	if e != nil {
		fmt.Print(e)
	}
	all_post := AllPostPage{
		Title: "Faqs Page",
		Faqs:  v.Faqs,
	}
	tmpl = template.Must(template.ParseFiles(tmpl_dir + "/index.html"))
	tmpl.Execute(w, all_post)
}

func faqAdmin(w http.ResponseWriter, r *http.Request) {
	v, e := getFagsJSON()
	if e != nil {
		fmt.Print(e)
	}
	faqs := AllPostPage{
		Title: "Admin Page",
		Faqs:  v.Faqs,
	}
	tmpl = template.Must(template.ParseFiles(tmpl_dir + "/admin.html"))
	tmpl.Execute(w, faqs)
}

func faqEdit(w http.ResponseWriter, r *http.Request) {
	validPath := regexp.MustCompile("^/faqs/(edit|save|view)/([a-zA-Z0-9]+)$")

	m := validPath.FindStringSubmatch(r.URL.Path)
	if m == nil {
		http.NotFound(w, r)
		return
	}
	faqId, err := strconv.Atoi(m[2])

	if err != nil {
		fmt.Print(err)
		return
	}

	v, e := getFagsJSON()
	if e != nil {
		fmt.Print(e)
	}

	var Faqs = v.Faqs
	var ff Faq

	for _, singleFaq := range Faqs {
		if singleFaq.ID == faqId {
			ff = singleFaq
		}
	}

	tmpl = template.Must(template.ParseFiles(tmpl_dir + "/edit.html"))
	page := EditPage{
		Title: "Edit Page",
		Faq:   ff,
	}
	tmpl.Execute(w, page)
}

func main() {
	port_address := ":9092"
	fmt.Println("Starting the todo app \n Should be accessible on port " + port_address)
	mux := http.NewServeMux()

	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	mux.HandleFunc("/faqs/all", allFaqs)
	mux.HandleFunc("/faqs/admin", faqAdmin)
	mux.HandleFunc("/faqs/edit/", faqEdit)

	log.Fatal(http.ListenAndServe(":9092", mux))
}
