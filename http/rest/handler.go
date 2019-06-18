package rest

func (a *App) NewRouter() {
	r := mux.NewRouter()
	r.HandleFunc("/", getHome).Methods("GET")
	r.HandleFunc("/customers", a.getCustomers).Methods("GET")
	r.HandleFunc("/customers", a.PostCustomers).Methods("POST")
	r.HandleFunc("/customers", a.PutCustomers).Methods("PUT").Queries("id", "{id}")
	r.HandleFunc("/customers", a.DeleteCustomers).Methods("DELETE").Queries("id", "{id}")
	a.router = r
}

func (a *App) getCustomers(w http.ResponseWriter, r *http.Request) {
	customers, err := ReadCustomers(a.db, 0)
	if err != nil {
		w.Header().Set("Content-type", "application/json")
		w.WriteHeader(http.StatusBadRequest)

		err := json.NewEncoder(w).Encode("Sorry, something bad happened.")
		if err != nil {
			log.Fatal(err)
		}
	}
	w.Header().Set("Content-type", "application/json")

	b, err := json.Marshal(customers)
	if err != nil {
		log.Fatal(err)
	}
	w.Write(b)
}
