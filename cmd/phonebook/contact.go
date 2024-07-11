package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
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

func (app *application) listContacts(w http.ResponseWriter, r *http.Request) {
	if rand.Intn(10) > 5 {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "se ferrou hein lek")
		app.counter500.Inc()
		return
	}

	resp, err := json.Marshal(map[string]any{"contacts": app.phonebook.contacts})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Server encountered a error... please try again later")
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}

func (app *application) listContactByID(w http.ResponseWriter, r *http.Request) {
	idParam := r.PathValue("id")
	id, err := strconv.Atoi(idParam)
	if err != nil || id < 0 {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "id must be a positive integer.")
		return
	}

	if id >= len(app.phonebook.contacts) {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "contact not found")
		return
	}

	resp, err := json.Marshal(app.phonebook.contacts[id])
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Server encountered a error... please try again later")
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}

func (app *application) createNewContact(w http.ResponseWriter, r *http.Request) {
	var ctt contact

	err := json.NewDecoder(r.Body).Decode(&ctt)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	app.phonebook.contacts = append(app.phonebook.contacts, ctt)
}
