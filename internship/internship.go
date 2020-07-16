package internship

// Positions holds neccessary information about each internship position
type Positions struct {
	CompanyName string
	Position    string
	Url         string
}

// Available is a container for internships found
var Available []Positions

// TODO: Move this to a config file
// Companies is a map of all target companies with their job listing url
var Companies = map[string]string{
	"Square":       "https://squareup.com/ca/careers/jobs?type=Intern",
	"Slack":        "https://slack.com/intl/en-ca/careers/university-recruiting#openings",
	"Stripe":       "https://stripe.com/jobs/search?s=intern",
	"Lyft":         "https://www.lyft.com/careers/university",
	"Uber":         "https://www.uber.com/us/en/careers/teams/university/",
	"Credit Karma": "https://www.creditkarma.com/careers/jobs/university",
	"Airbnb":       "https://careers.airbnb.com/university/",
	"Pinterest":    "https://www.pinterestcareers.com/university",
	"Mozilla":      "https://careers.mozilla.org/listings/?position_type=Intern",
	"Palantir":     "https://www.palantir.com/careers/",
}

// Keyword holds the keywords to parse for in each listing
const Keyword string = "mobile"
