package hackerone

const (
	// BaseAPI is the base URL for the HackerOne API
	BaseAPI = "https://api.hackerone.com/"

	// ProgramHandlesAPIUrl is the base URL of the HackerOne API for scraping program handles
	ProgramHandlesAPIUrl = BaseAPI + "v1/hackers/programs?page[size]=100&page[number]="

	// ProgramHandleScopesAPIUrl is the api URL for scraping structured scope or assets of a program
	//
	// must replace {handle} with the target program handle to be scraped
	ProgramHandleScopesAPIUrl = BaseAPI + "v1/hackers/programs/{handle}/structured_scopes"
)
