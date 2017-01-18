# POST /api/patients

  {
    "name": {
      "given": "Emily"
      "family": "Smith"
    }
    "address": {
      "streetAddressLine": "1212 1st Street"
      "city": "Ames"
      "state": "IA"
      "postalCode": "50011"
    }
    "gender": "Female"
    "birthTime": "1995-07-14" // YYYY-MM-DD
    "ssn": "123-45-6789"
    "allergies": [
      {
        "pid": 1
        "title": "Something"
        "Reaction": "Something"
        "Comments": "Some text"
        "Occurrence": 1
        "Severity": "Mild"
        "Outcome": 1
      }
    ]
    "medicalIssues": [
      {
        "pid": 1
        "title": "Something"
        "Comments": "Some text"
        "Occurrence": 1
        "Severity": "Mild"
        "Outcome": 1
      }
    ]
    "medications": [
      {
        "pid": 1
        "title": "Something"
        "Comments": "Some text"
        "Occurrence": 1
        "Severity": "Mild"
        "Outcome": 1
      }
    ]
    "extension": {
      "username": "emilysmith95"
      "password": "ilovecats"
    }
  }

## CURL

    curl http://localhost:8000/api/client -X POST -H "Content-Type: application/json"  -d '{"name": {"given": "Emily", "family": "Smith"}, "address": {"streetAddressLine": "1212 1st Street", "city": "Ames", "state": "IA", "postalCode": "50011"}, "gender": "Female", "birthTime": "1995-07-14", "ssn": "123-45-6789", "extension": {"username": "emilysmith95", "password": "ilovecats"}}'
