package db

// defines the structure of the service fields
type Service struct {
	ServiceID          int
	ServiceName        string
	ServiceDescription string
	ServiceLogo        string
	ServiceWebsite     string
	Categories         []Category
}

// defines the structure of the category fields
type Category struct {
	CategoryID          int
	CategoryTitle       string
	CategoryDescription string
	CategoryKey         string
}
