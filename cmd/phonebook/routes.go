package main

func (app *application) bindRoutes() {
	app.mux.HandleFunc("GET /list", app.listContacts)
	app.mux.HandleFunc("GET /list/{id}", app.listContactByID)
	app.mux.HandleFunc("POST /", app.createNewContact)
}
