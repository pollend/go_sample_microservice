package controller

import (
    "github.com/gorilla/mux"
    "net/http"
    "prospect/mobile-physician-api/pkg/repositories"
)

type PetController struct {
    provider* repositories.ProviderRepository
}

func (controller PetController) Route(router* mux.Router) {
    router.HandleFunc("/pet", controller.index)
    router.HandleFunc("/pet/{id}", controller.get)
}

func (Controller PetController) index(w http.ResponseWriter, r *http.Request)  {
}

func (Controller PetController) get(w http.ResponseWriter, r *http.Request) {

}
