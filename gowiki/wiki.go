package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// Page represents the structure for a wiki page
type Page struct {
	Title string
	Body  []byte
}

func (p *Page) save() error {
	filename := p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/view/"):]
	p, _ := loadPage(title)
	fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", p.Title, p.Body)
}

func mock() {
	fmt.Println("Mock IO")
	p1 := &Page{Title: "test", Body: []byte("This is a sample Page.")}
	p1.save()
	p2, _ := loadPage("test")
	fmt.Println(string(p2.Body))
}

func main() {
	mock()
	http.HandleFunc("/view/", viewHandler)
	log.Fatal(http.ListenAndServe(":0880", nil))
}
