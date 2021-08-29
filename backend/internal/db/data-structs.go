package db

// defines the structure of the service fields

type Service struct {
	ServiceID          int
	ServiceName        string
	ServiceDescription string
	ServiceLogo        string
	ServiceWebsite     string
	ServiceHash        string
	ServiceInserted    int
	ServiceCreator     string
	ServiceContact     string
	CategoryKeys       []string
}

// defines the structure of the category fields
type Category struct {
	CategoryID          int
	CategoryTitle       string
	CategoryDescription string
	CategoryKey         string
	ServiceCount        int
}
