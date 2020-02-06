package controller

import (
    "github.com/gorilla/mux"
    "net/http"
)

type PetStoreController struct {
}

func (controller PetStoreController) Route(router* mux.Router) {
    router.HandleFunc("/store", controller.index)
    router.HandleFunc("/store/{id}", controller.get)
}

func (Controller PetStoreController) index(w http.ResponseWriter, r *http.Request)  {
}

func (Controller PetStoreController) get(w http.ResponseWriter, r *http.Request) {

}
