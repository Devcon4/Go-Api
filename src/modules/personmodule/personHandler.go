package personmodule

import (
	"net/http"

	"github.com/Devcon4/Go-Api/tools/app"
	"github.com/gorilla/mux"
)

// PersonHandler : Controller for persons
type PersonHandler struct {
	router        *mux.Router
	personService *PersonService
}

// NewPersonHandler : Creates a PersonHandler
func NewPersonHandler(router *mux.Router, personService *PersonService) *PersonHandler {
	return &PersonHandler{router, personService}
}

// Register : Wires up route handlers
func (h PersonHandler) Register() {
	h.router.HandleFunc("/person", h.Get).Methods("GET")
	h.router.HandleFunc("/persons", h.GetList).Methods("GET")
}

// Get : Get handler func
func (h PersonHandler) Get(w http.ResponseWriter, r *http.Request) {
	u, err := h.personService.Get()
	app.JSONHandler(w, u, err)
}

// GetList : GetList handler func
func (h PersonHandler) GetList(w http.ResponseWriter, r *http.Request) {
	l, err := h.personService.GetList()
	app.JSONHandler(w, l, err)
}
