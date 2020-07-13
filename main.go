package main

import (
	"fmt"
	"io"
	"os"
	"regexp"
	"sync"

	"github.com/PuerkitoBio/goquery"
	"github.com/rkabani19/internship-assistant/client"
	"github.com/rkabani19/internship-assistant/internship"
)

type intershipPositions struct {
	companyName string
	position    string
	url         string
}

var (
	internshipsAvailable []intershipPositions
	mutex                sync.Mutex
)

func getInternshipLinks(company string, body io.ReadCloser) {
	doc, err := goquery.NewDocumentFromReader(body)
	if err != nil {
		return
	}

	doc.Find("a[href]").Each(func(index int, item *goquery.Selection) {
		href, _ := item.Attr("href")
		title := item.Text()

		match, _ := regexp.MatchString(fmt.Sprintf(`(?i)%s\b`, internship.Keyword), title)
		if match {
			mutex.Lock()
			internshipsAvailable = append(internshipsAvailable, intershipPositions{
				companyName: company,
				position:    title,
				url:         href,
			})
			mutex.Unlock()
		}
	})
}

func internshipWorker(company string, url string, wg *sync.WaitGroup) {
	internshipClient := client.NewInternshipClient(url)
	resp, err := internshipClient.Fetch()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
	defer resp.Body.Close()
	defer wg.Done()

	getInternshipLinks(company, resp.Body)
}

func printInternships() {
	var previousCompany string
	for _, internshipAvailable := range internshipsAvailable {
		if previousCompany == "" || previousCompany != internshipAvailable.companyName {
			fmt.Printf("%s:\n\t%s - %s\n", internshipAvailable.companyName, internshipAvailable.position, internshipAvailable.url)
		} else {
			fmt.Printf("\t%s - %s\n", internshipAvailable.position, internshipAvailable.url)
		}
		previousCompany = internshipAvailable.companyName
	}
}

func main() {
	var wg sync.WaitGroup

	for company, url := range internship.Companies {
		wg.Add(1)
		go internshipWorker(company, url, &wg)
	}

	wg.Wait()
	printInternships()
}
