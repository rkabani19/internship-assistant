package main

import (
	"fmt"
	"io"
	"os"
	"regexp"

	"github.com/PuerkitoBio/goquery"
	"github.com/rkabani19/internship-assistant/client"
	"github.com/rkabani19/internship-assistant/internship"
)

type intershipPositions struct {
	companyName string
	position    string
	url         string
}

var internshipsAvailable []intershipPositions

func main() {
	for company, url := range internship.Companies {
		internshipClient := client.NewInternshipClient(url)
		resp, err := internshipClient.Fetch()
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
			os.Exit(1)
		}
		defer resp.Body.Close()

		fmt.Printf("%s\n", company)
		getInternshipLinks(resp.Body)
	}
}

func getInternshipLinks(body io.ReadCloser) {
	doc, err := goquery.NewDocumentFromReader(body)
	if err != nil {
		return
	}

	doc.Find("a[href]").Each(func(index int, item *goquery.Selection) {
		href, _ := item.Attr("href")
		title := item.Text()

		match, _ := regexp.MatchString(fmt.Sprintf(`(?i)%s\b`, internship.Keyword), title)
		if match {
			fmt.Printf("link: %s - anchor text: %s\n", href, item.Text())
		}
	})
}
