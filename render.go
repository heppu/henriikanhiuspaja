package main

import (
	"encoding/json"
	"html/template"
	"log"
	"os"
)

type Data struct {
	Phone     string `json:"phone"`
	Catalogue []struct {
		Header string `json:"header"`
		Items  []struct {
			Name  string `json:"name"`
			Price string `json:"price"`
		} `json:"items"`
	} `json:"catalogue"`
}

func main() {
	tmpl, err := template.ParseFiles("index.tmpl")
	if err != nil {
		log.Fatal("Could not parse template:", err)
	}

	input, err := os.Open("data.json")
	if err != nil {
		log.Fatal("Could not open data:", err)
	}
	defer input.Close()

	data := &Data{}
	if err := json.NewDecoder(input).Decode(data); err != nil {
		log.Fatal("Could not decode data:", err)
	}

	output, err := os.Create("dist/index.html")
	if err != nil {
		log.Fatal("Could not parse template:", err)
	}
	defer output.Close()

	if err := tmpl.Execute(output, data); err != nil {
		log.Fatal("Could not render template:", err)
	}

	log.Println("Render complete")
}
