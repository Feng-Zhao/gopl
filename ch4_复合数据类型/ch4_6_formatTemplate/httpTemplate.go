package main

import (
	"fmt"
	"html/template"
	"os"
)

const httpTemplate = `
	<h1>{{.Title}}</h1>
	<p>{{.Content}}</p>
	`

type HtmlData struct {
	Title   string
	Content string
}

// go run httpTemplate.go > 1.html
func main() {
	data := HtmlData{Title: "title", Content: "this is a html file composed from template"}
	tmp, err := template.New("html").Parse(httpTemplate)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	tmp.Execute(os.Stdout, data)
}
