package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gofrs/uuid"
	"github.com/gorilla/mux"
	"github.com/stackpath/backend-developer-tests/rest-service/pkg/models"
)

func main() {
	fmt.Println("SP// Backend Developer Test - RESTful Service")
	fmt.Println()

	r := mux.NewRouter()

	// Routes:
	r.Path("/people").Queries("phone_number", "{phone_number}").HandlerFunc(GetPeopleFilterPhone).Methods("GET")
	r.Path("/people").Queries("first_name", "{first_name}", "last_name", "{last_name}").HandlerFunc(GetPeopleFilterName).Methods("GET")
	r.HandleFunc("/people", GetAllPeople).Methods("GET")
	r.HandleFunc("/people/{id}", GetPerson).Methods("GET")

	port := ":9000"
	fmt.Printf("Serving at %s \n", port)
	if err := http.ListenAndServe(port, r); err != nil {
		log.Fatal(err)
	}
}

//GetAllPeople responses with all the peoples in the system
//	@Success 200
func GetAllPeople(w http.ResponseWriter, r *http.Request) {
	log.Println("Returning all people")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(models.AllPeople())
	return
}

//GetPerson responses with info of person
//	@Success 200
//	@Not Found 404 id person doesnt exist
func GetPerson(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	log.Printf("Retreiving Person %v \n", id)
	p, err := models.FindPersonByID(uuid.FromStringOrNil(id))
	if err != nil {
		w.WriteHeader(404)
		return
	}
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(p)
	return
}

//GetPeopleFilterName responses with people that match the first and last name
//	@Success 200
func GetPeopleFilterName(w http.ResponseWriter, r *http.Request) {
	firstName := r.FormValue("first_name")
	lastName := r.FormValue("last_name")
	log.Printf("Retreiving People with first-name: %s and last-name: %s \n", firstName, lastName)
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(models.FindPeopleByName(firstName, lastName))
	return
}

//GetPeopleFilterPhone responses with people that match the phone_number
//	@Success 200
func GetPeopleFilterPhone(w http.ResponseWriter, r *http.Request) {
	phone := r.FormValue("phone_number")
	fmt.Println(phone)
	log.Printf("Retreiving People with phone-number : %v \n", phone)
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(models.FindPeopleByPhoneNumber(phone))
	return
}
