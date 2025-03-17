package main

import (
	"encoding/json"
	"fmt"
	"os"
	"text/template"
)

// Photo represents the structure of each photo entry
type Photo struct {
	Title       string `json:"title"`
	Date        string `json:"date"`
	Dimensions  []int  `json:"dimensions"`
	Description string `json:"description"`
	Asset       string `json:"asset"`
	Portrait    bool   `json:"portrait"`
	Alt         string `json:"alt"`
}

// Data represents the root structure
type PhotographyData struct {
	Data []Photo `json:"data"`
}

func parse() (PhotographyData, error) {
	// Read the JSON file
	file, err := os.ReadFile("Photography.json")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return PhotographyData{}, err
	}

	var data PhotographyData

	// Parse JSON data
	err = json.Unmarshal(file, &data)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return PhotographyData{}, err
	}

	return data, nil
}

func main() {
	tmpl, err := template.ParseFiles("template.html")
	if err != nil {
		fmt.Println("Error parsing template:", err)
		return
	}

	data, err := parse()
	if err != nil {
		fmt.Println("Error parsing data:", err)
		return
	}

	outputFile, err := os.Create("index.html")
	if err != nil {
		fmt.Println("Error creating output file:", err)
		return
	}

	defer outputFile.Close()

	err = tmpl.Execute(outputFile, data)
	if err != nil {
		fmt.Println("Error executing template:", err)
		return
	}

	fmt.Println("HTML file generated successfully.")
}
