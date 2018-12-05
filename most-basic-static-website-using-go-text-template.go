package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

const masterTemplate = `<!DOCTYPE html>
<html>
  <head>
    <meta charset="utf-8">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <title>{{.Title}}</title>
	<link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/bulma/0.7.2/css/bulma.min.css">
	{{.Header}}
  </head>
  <body>
  <nav class="navbar" role="navigation" aria-label="main navigation">
  <a class="navbar-item" href="/">
        Home
  </a>
  <a class="navbar-item" href="/about">
        About
  </a>
</nav>
  <section class="section">
    <div class="container">
      <h1 class="title">
	  {{.Title}}
      </h1>
      <div class="hero-body">
	  {{.Body}}
      </div>
    </div>
  </section>
  {{.Footer}}
  </body>
</html>`

type pageData struct {
	Title, Header, Body, Footer string
}

func handleHome(res http.ResponseWriter, req *http.Request) {
	homePageData := pageData{
		Title: "My home page title",
		Body:  "<p>My home page contents...</p>"}

	homePage := template.New("home")
	homePage, err := homePage.Parse(masterTemplate)

	if err != nil {
		log.Fatal("Failed to parse template: ", err)
		return
	}

	res.Header().Set("content-type", "text/html; charset=utf-8")

	err = homePage.Execute(res, homePageData)
	if err != nil {
		log.Fatal("Failed to execute template: ", err)
		return
	}
}

func handleAbout(res http.ResponseWriter, req *http.Request) {
	homePageData := pageData{
		Title: "My about page title",
		Body:  "<p>My about page contents...</p>"}

	homePage := template.New("home")
	homePage, err := homePage.Parse(masterTemplate)

	if err != nil {
		log.Fatal("Failed to parse template: ", err)
		return
	}

	res.Header().Set("content-type", "text/html; charset=utf-8")

	err = homePage.Execute(res, homePageData)
	if err != nil {
		log.Fatal("Failed to execute template: ", err)
		return
	}
}

func main() {
	http.HandleFunc("/", handleHome)
	http.HandleFunc("/about", handleAbout)

	fmt.Println("Server started at http://localhost:8000")

	if err := http.ListenAndServe(":8000", nil); err != nil {
		panic(err)
	}
}
