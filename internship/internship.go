package internship

// Positions holds neccessary information about each internship position
type Positions struct {
	CompanyName string
	Position    string
	Url         string
}

// Available is a container for internships found
var Available []Positions

// Company holds information about internhship search point
type Company struct {
	Company string
	Url     string
}

// Keyword holds the keywords to parse for in each listing
const Keyword string = "intern" // TODO: allow user to set
