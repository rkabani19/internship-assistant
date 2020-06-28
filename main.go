package main

import (
	"fmt"
	"io"
	"os"
	"strings"

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
	internshipClient := client.NewInternshipClient("https://careers.mozilla.org/listings/")
	resp, err := internshipClient.Fetch()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	getInternshipLinks(resp.Body)
}

func getInternshipLinks(body io.ReadCloser) {
	doc, err := goquery.NewDocumentFromReader(body)
	if err != nil {
		return
	}

	doc.Find("a[href]").Each(func(index int, item *goquery.Selection) {
		href, _ := item.Attr("href")
		title := item.Text()
		if strings.Contains(strings.ToLower(title), strings.ToLower(internship.Keyword)) {
			fmt.Printf("keyword: %s - link: %s - anchor text: %s\n", internship.Keyword, href, item.Text())
		}
	})
}

// TODO: Iterate through map of links
