package main

import (
	"log"
	"net/http"
	"path/filepath"
	"sync"
	"text/template"
)

type templateHandler struct {
	once     sync.Once
	filename string
	templ    *template.Template
}

// SeveHTTP: handle http request
func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// It reads template file at once
	t.once.Do(func() {
		t.templ = template.Must(template.ParseFiles(filepath.Join("templates", t.filename)))
	})
	if err := t.templ.Execute(w, nil); err != nil {
		log.Fatal("ServeHTTP:", err)
	}
}

func main() {
	// root
	http.Handle("/", &templateHandler{filename: "chat.html"})
	// it starts web server
	if err := http.ListenAndServe(":18888", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
