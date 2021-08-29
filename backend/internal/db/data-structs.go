package db

// defines the structure of the service fields
type Service struct {
	ServiceID          int
	ServiceName        string
	ServiceDescription string
	ServiceLogo        string
	ServiceWebsite     string
	/*
		Hash consists of ServiceName, ServiceWebsite and ServiceInserted
	*/
	ServiceHash     string
	ServiceInserted int
	ServiceCreator  string
	//Categories      []Category
	CategoryKeys []string
}

// defines the structure of the category fields
type Category struct {
	CategoryID          int
	CategoryTitle       string
	CategoryDescription string
	CategoryKey         string
	ServiceCount        int
}
