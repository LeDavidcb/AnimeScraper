package main

import (
	"flag"
	"fmt"
	"strings"

	"github.com/gocolly/colly"
	"github.com/nyaaScrapper/pkg"
)

func main() {
	// Get flag
	nameflag := flag.String("name", "", "Name of the anime you want to look up")
	flag.Parse()
	var name string = *nameflag

	if name == "" {
		fmt.Println(`No anime name was provided, provide the anime name by using the flag name: --name=<url>`)
		return
	}

	c := colly.NewCollector()
	fmt.Printf("Searching for anime %v \n", name)

	entries, err := pkg.ReturnEntries(c, fmt.Sprintf("https://nyaa.si/?f=0&c=0_0&q=%v", name))
	if err != nil {
		fmt.Println("Error at scraping entities, more info: ", err)
		return
	}

	fmt.Println("Found", len(entries), "Animes")
	fmt.Println("Do you want to display them? [y/n]")

	var response string
	_, err = fmt.Scan(&response)
	if err != nil {
		fmt.Println("Error while getting response:", err)
	}
	switch strings.Trim(strings.ToLower(response), " ") {
	case "y":
		entries.Display()
		response = "" // reset response
	case "n":
	default:
		fmt.Println("Not a valid response")
	}

	fmt.Println("Do you want to coppy all magnets? [y/n]")
	_, err = fmt.Scan(&response)
	if err != nil {
		fmt.Println("Error while getting response:", err)
	}
	switch strings.Trim(strings.ToLower(response), " ") {
	case "y":
		err = entries.CoppyMagnets()
		if err != nil {
			fmt.Println(err)
		}
	case "n":
	default:
		fmt.Println("Not a valid response")
	}
}
