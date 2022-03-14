package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

func init(){
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main(){
	http.HandleFunc("/", hom)
	http.HandleFunc("/about", abo)
	http.HandleFunc("/test", tes)
	http.HandleFunc("/test1", tes1)
	http.Handle("/stuff/", http.StripPrefix("/stuff", http.FileServer(http.Dir("./assets"))))
	http.ListenAndServeTLS(":8081", "server.crt", "server.key", nil)
}

func hom(w http.ResponseWriter, r *http.Request){
	if pusher, ok := w.(http.Pusher); ok {
		// Push is supported.
		if err := pusher.Push("/stuff/favicon.ico", nil); err != nil {
			log.Printf("Failed to push: %v", err)
		}
	}
	tpl.ExecuteTemplate(w, "default.gohtml", nil)
}

func abo(w http.ResponseWriter, r *http.Request){
	if pusher, ok := w.(http.Pusher); ok {
		// Push is supported.
		for i, j, k := 1, 1, 1; i < 101; i, j, k = i+1, i/10+1, i%10+1 {
			target := fmt.Sprintf("/stuff/row-%d-column-%d.jpg", j, k)
			if err := pusher.Push(target, nil); err != nil {
				log.Printf("Failed to push: %v", err)
			}
		}
		if err := pusher.Push("/stuff/favicon.ico", nil); err != nil {
			log.Printf("Failed to push: %v", err)
		}
	}
	tpl.ExecuteTemplate(w, "about.gohtml", nil)
}

func tes(w http.ResponseWriter, r *http.Request){
	if pusher, ok := w.(http.Pusher); ok {
		// Push is supported.
		for i, j, k := 1, 1, 1; i < 101; i, j, k = i+1, i/10+1, i%10+1 {
			target := fmt.Sprintf("/stuff/row-%d-column-%d.jpg", j, k)
			if err := pusher.Push(target, nil); err != nil {
				log.Printf("Failed to push: %v", err)
			}
		}
		if err := pusher.Push("/stuff/favicon.ico", nil); err != nil {
			log.Printf("Failed to push: %v", err)
		}
	}
	tpl.ExecuteTemplate(w, "test.gohtml", nil)
}

func tes1(w http.ResponseWriter, r *http.Request)  {
	tpl.ExecuteTemplate(w, "test1.gohtml", nil)
}