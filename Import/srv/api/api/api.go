package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"os/exec"
	"strconv"
)

type ClientHandler struct {
	Input ClientHandlerRequest
}

func (c *ClientHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	headers(w, r)

	switch r.Method {
	case "OPTIONS":
		break
	case "POST":
		c.respondPOST(w, r)
		break
	default:
		respondNotImplemented(w, r)
	}
}

func (c *ClientHandler) respondPOST(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&c.Input)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("{\"status\": \"error\", \"message\": \"Error parsing JSON request data.\"}"))
		return
	}

	pid, err := c.CreateClient(w, r)
	if err != nil {
		return
	}
	c.CreateCreds(w, r)
	c.CreateAllergies(w, r, pid)
	c.CreateMedicalIssue(w, r, pid)
	c.CreateMedications(w, r, pid)
}

func (c *ClientHandler) CreateClient(w http.ResponseWriter, r *http.Request) (int, error) {
	var err error
	var isemr_create ClientHandlerISEmrCreate
	var isemr_create_byte []byte
	isemr_create.FirstName = c.Input.Name.Given
	isemr_create.LastName = c.Input.Name.Family
	isemr_create.Street = c.Input.Address.StreetAddressLine
	isemr_create.City = c.Input.Address.City
	isemr_create.State = c.Input.Address.State
	isemr_create.PostalCode = c.Input.Address.PostalCode
	isemr_create.Sex = c.Input.Gender
	isemr_create.SSN = c.Input.SSN
	isemr_create.DOB = c.Input.BirthTime

	isemr_create_byte, err = json.Marshal(isemr_create)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("{\"status\": \"error\", \"message\": \"Error generating new JSON from parsed JSON.\"}"))
		return -1, err
	}

	var url string = isemr_addr + "/api/create_patient.php"
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(isemr_create_byte))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("{\"status\": \"error\", \"message\": \"Error creating new client.\"}"))
		return -1, err
	}
	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	buf := bytes.NewBuffer(make([]byte, 0, resp.ContentLength))
	_, _ = buf.ReadFrom(resp.Body)
	body := buf.Bytes()

	str := string(body)
	log.Println(str)
	pid, err := strconv.Atoi(str)

	return pid, err
}

func (c *ClientHandler) CreateCreds(w http.ResponseWriter, r *http.Request) {
	c.CreateCredsISEmr(w, r)
	c.CreateCredsPatientView(w, r)
}

func (c *ClientHandler) CreateCredsISEmr(w http.ResponseWriter, r *http.Request) {
	var err error
	var isemr_creds ClientHandlerISEmrCredentials
	var isemr_creds_byte []byte
	isemr_creds.SSN = c.Input.SSN
	isemr_creds.Username = c.Input.Extension.Username
	isemr_creds.Password = c.Input.Extension.Password

	isemr_creds_byte, err = json.Marshal(isemr_creds)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("{\"status\": \"error\", \"message\": \"Error generating new JSON from parsed JSON.\"}"))
	}

	var url string = isemr_addr + "/api/set_credentials.php"
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(isemr_creds_byte))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("{\"status\": \"error\", \"message\": \"Error creating new client credentials.\"}"))
	}
	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}
	log.Println(resp.Body)
}

func (c *ClientHandler) CreateCredsPatientView(w http.ResponseWriter, r *http.Request) {
	// Assume that we are running on the patient portal
	cmd := exec.Command("bash", "-c", "useradd -g patients -m -p $(echo \""+c.Input.Extension.Password+"\" | openssl passwd -1 -stdin) "+c.Input.Extension.Username)
	cmd.Run()
}

func (c *ClientHandler) CreateAllergies(w http.ResponseWriter, r *http.Request, pid int) {
	for i := range c.Input.Allergies {
		go c.CreateAllergy(w, r, pid, i)
	}
}

func (c *ClientHandler) CreateAllergy(w http.ResponseWriter, r *http.Request, pid int, i int) {
	var err error
	a := c.Input.Allergies[i]
	a.PID = pid

	var isemr_allergy_byte []byte
	isemr_allergy_byte, err = json.Marshal(a)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("{\"status\": \"error\", \"message\": \"Error generating new JSON from parsed JSON.\"}"))
	}

	var url string = isemr_addr + "/api/create_allergy.php"
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(isemr_allergy_byte))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("{\"status\": \"error\", \"message\": \"Error creating new client allergy.\"}"))
	}
	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	buf := bytes.NewBuffer(make([]byte, 0, resp.ContentLength))
	_, _ = buf.ReadFrom(resp.Body)
	body := buf.Bytes()

	str := string(body)
	log.Println("CreateAllergies")
	log.Println(resp.Status)
	log.Println(resp.StatusCode)
	log.Println(string(str))
}

func (c *ClientHandler) CreateMedicalIssue(w http.ResponseWriter, r *http.Request, pid int) {
	for i := range c.Input.MedicalIssues {
		go c.CreateMedicalIssueSingle(w, r, pid, i)
	}
}

func (c *ClientHandler) CreateMedicalIssueSingle(w http.ResponseWriter, r *http.Request, pid int, i int) {
	var err error
	m := c.Input.MedicalIssues[i]
	m.PID = pid

	var isemr_medical_issues_byte []byte
	isemr_medical_issues_byte, err = json.Marshal(m)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("{\"status\": \"error\", \"message\": \"Error generating new JSON from parsed JSON.\"}"))
	}

	var url string = isemr_addr + "/api/create_medical_issues.php"
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(isemr_medical_issues_byte))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("{\"status\": \"error\", \"message\": \"Error creating new client medical_issues.\"}"))
	}
	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	buf := bytes.NewBuffer(make([]byte, 0, resp.ContentLength))
	_, _ = buf.ReadFrom(resp.Body)
	body := buf.Bytes()

	str := string(body)
	log.Println("CreateMedicalIssues")
	log.Println(resp.Status)
	log.Println(resp.StatusCode)
	log.Println(string(str))
}

func (c *ClientHandler) CreateMedications(w http.ResponseWriter, r *http.Request, pid int) {
	for i := range c.Input.Medications {
		go c.CreateMedication(w, r, pid, i)
	}
}

func (c *ClientHandler) CreateMedication(w http.ResponseWriter, r *http.Request, pid int, i int) {
	var err error
	m := c.Input.Medications[i]
	m.PID = pid

	var isemr_medication_byte []byte
	isemr_medication_byte, err = json.Marshal(m)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("{\"status\": \"error\", \"message\": \"Error generating new JSON from parsed JSON.\"}"))
	}

	log.Println(string(isemr_medication_byte))

	var url string = isemr_addr + "/api/create_medication.php"
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(isemr_medication_byte))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("{\"status\": \"error\", \"message\": \"Error creating new client medication.\"}"))
	}
	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	buf := bytes.NewBuffer(make([]byte, 0, resp.ContentLength))
	_, _ = buf.ReadFrom(resp.Body)
	body := buf.Bytes()

	log.Println("CreateMedications")
	log.Println(resp.Status)
	log.Println(resp.StatusCode)
	log.Println(string(body))
}
