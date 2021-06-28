package hocontrollers

import (
	models "WebDevelopmentTodd/MongoDB/MongoDB/HandsOn/homodels"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	uuid "github.com/satori/go.uuid"
)

type u_map map[string]models.User

func NewUserMap() *u_map {
	myMap := u_map{}
	return &myMap
}

func (um u_map) GetUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// Grab id
	id := p.ByName("id")

	// composite literal
	u := models.User{}

	u, ok := um[id]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// Marshal provided interface into JSON structure
	uj, _ := json.Marshal(u)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK) // 200
	fmt.Fprintf(w, "%s\n", uj)
}

func (um u_map) CreateUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Println("Create User")
	u := models.User{}

	json.NewDecoder(r.Body).Decode(&u)

	uui, _ := uuid.NewV4()
	u.Id = uui.String()	
	um[u.Id] = u
	fmt.Println("Failed here")

	uj, _ := json.Marshal(u)

	fmt.Println(um)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated) // 201
	fmt.Fprintf(w, "%s\n", uj)
}

func (um u_map) DeleteUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")
	u := models.User{}

	u, ok := um[id]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
	}
	delete(um, id)
	uj, _ := json.Marshal(u)
	w.WriteHeader(http.StatusOK) // 200
	fmt.Fprint(w, "Deleted user", uj, "\n")
}
