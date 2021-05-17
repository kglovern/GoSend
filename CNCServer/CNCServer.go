package CNCServer

import (
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"log"
	"net/http"
)

type CNCServer struct {
	Router *mux.Router
	DB     *gorm.DB
}

func (s *CNCServer) Initialize() {
	s.Router = mux.NewRouter()
	s.initializeRoutes()
}

func (s *CNCServer) initializeRoutes() {
	s.Router.HandleFunc("/", s.getHomePage)
}

func (s *CNCServer) Run(port string) {
	log.Fatal(http.ListenAndServe(port, s.Router))
}
