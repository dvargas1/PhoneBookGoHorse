package main

import (
	"bytes"
	"text/template"
)

type contact struct {
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Description string `json:"description"`
	PhoneNumber string `json:"phone_number"`
	Age         int    `json:"age"`
}

func (c contact) String() string {
	var buf bytes.Buffer

	templateString := `
=================================
FirstName:   {{ .FirstName }}
LastName:    {{ .LastName }}
Age:         {{ .Age }}
Description: {{ .Description }}
=================================
    `

	tmpl := template.Must(template.New("contact").Parse(templateString))

	if err := tmpl.Execute(&buf, c); err != nil {
		panic("failed to parse correctly the contact")
	}

	return buf.String()
}
