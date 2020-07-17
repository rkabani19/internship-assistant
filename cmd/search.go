package cmd

import (
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
	"sync"

	"github.com/PuerkitoBio/goquery"
	"github.com/rkabani19/internship-assistant/client"
	"github.com/rkabani19/internship-assistant/internship"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// searchCmd represents the search command
var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "Search for internships",
	Run: func(cmd *cobra.Command, args []string) {
		search()
	},
}

var mutex sync.Mutex

func init() {
	rootCmd.AddCommand(searchCmd)
}

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
			internship.Available = append(internship.Available, internship.Positions{
				CompanyName: company,
				Position:    title,
				Url:         href,
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
	for _, internshipAvailable := range internship.Available {
		if previousCompany == "" || previousCompany != internshipAvailable.CompanyName {
			fmt.Printf("%s:\n\t%s - %s\n", strings.ToUpper(internshipAvailable.CompanyName), internshipAvailable.Position, internshipAvailable.Url)
		} else {
			fmt.Printf("\t%s - %s\n", internshipAvailable.Position, internshipAvailable.Url)
		}
		previousCompany = internshipAvailable.CompanyName
	}
}

func search() {
	var wg sync.WaitGroup

	fmt.Println("Fetching internships...")
	for company, url := range viper.GetStringMapString("companies") {
		wg.Add(1)
		go internshipWorker(company, url, &wg)
	}

	wg.Wait()
	printInternships()
}
