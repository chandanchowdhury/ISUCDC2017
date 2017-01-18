package main

type ClientHandlerRequest struct {
	Name          ClientHandlerRequestName            `json:"name"`
	Address       ClientHandlerRequestAddress         `json:"address"`
	Gender        string                              `json:"gender"`
	BirthTime     string                              `json:"birthTime"`
	SSN           string                              `json:"ssn"`
	Allergies     []ClientHandlerRequestAllergy       `json:"allergies"`
	MedicalIssues []ClientHandlerRequestMedicalIssues `json:"medicalIssues"`
	Medications   []ClientHandlerRequestMedications   `json:"medications"`
	Extension     ClientHandlerRequestExtension       `json:"extension"`
}

type ClientHandlerRequestName struct {
	Given  string `json:"given"`
	Family string `json:"family"`
}

type ClientHandlerRequestAddress struct {
	StreetAddressLine string `json:"streetAddressLine"`
	City              string `json:"city"`
	State             string `json:"state"`
	PostalCode        string `json:"postalCode"`
}

type ClientHandlerRequestExtension struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type ClientHandlerRequestAllergy struct {
	PID        int    `json:"pid"`
	Title      string `json:"title"`
	Reaction   string `json:"Reaction"`
	Comments   string `json:"Comments"`
	Occurrence int    `json:"Occurrence"`
	Severity   string `json:"Severity"`
	Outcome    int    `json:"Outcoe"`
}

type ClientHandlerRequestMedicalIssues struct {
	PID        int    `json:"pid"`
	Title      string `json:"title"`
	Comments   string `json:"Comments"`
	Occurrence int    `json:"Occurrence"`
	Severity   string `json:"Severity"`
	Outcome    int    `json:"Outcoe"`
}

type ClientHandlerRequestMedications struct {
	PID        int    `json:"pid"`
	Title      string `json:"title"`
	Comments   string `json:"Comments"`
	Occurrence int    `json:"Occurrence"`
	Severity   string `json:"Severity"`
	Outcome    int    `json:"Outcoe"`
}

type ClientHandlerISEmrCreate struct {
	FirstName  string `json:"fname"`
	LastName   string `json:"lname"`
	Street     string `json:"street"`
	City       string `json:"city"`
	State      string `json:"state"`
	PostalCode string `json:"postal_code"`
	Sex        string `json:"sex"`
	SSN        string `json:"ss"`
	DOB        string `json:"DOB"`
}

type ClientHandlerISEmrCredentials struct {
	SSN      string `json:"SSN"`
	Username string `json:"Username"`
	Password string `json:"Password"`
}
