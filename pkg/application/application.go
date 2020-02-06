package application

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"prospect/mobile-physician-api/pkg/controller"
)

const (
	STATIC_DIR = "/public/"
	PORT = "8000"
)

func NewRouter(petStore *controller.PetStoreController, pet *controller.PetController) *mux.Router {
	mux := mux.NewRouter()
	mux.PathPrefix(STATIC_DIR).Handler(http.StripPrefix(STATIC_DIR,http.FileServer(http.Dir("." + STATIC_DIR))))

	// setup api routes
    api := mux.PathPrefix("/api/").Subrouter()
    petStore.Route(api)
	pet.Route(api)
	return mux
}


type Application struct {
	router *mux.Router
}

func NewApplication(router *mux.Router) Application {
	return Application{
		router: router,
	}
}

func (app Application) Start() {
	srv := &http.Server{
		Addr:    ":" + PORT,
		Handler: app.router,
	}
	log.Print("Starting server on port :8000")
	log.Fatal(srv.ListenAndServe())
}


