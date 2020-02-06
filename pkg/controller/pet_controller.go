package controller

import (
    "github.com/gorilla/mux"
    "net/http"
)

type PetController struct {
}

func NewPetController() *PetController {
    return &PetController{
    }
}

func (controller PetController) Route(router* mux.Router) {
    router.HandleFunc("/pet", controller.index)
    router.HandleFunc("/pet/{id}", controller.get)
}

func (Controller PetController) index(w http.ResponseWriter, r *http.Request)  {
}

func (Controller PetController) get(w http.ResponseWriter, r *http.Request) {

}
