package CNCServer

import (
	"github.com/gorilla/mux"
	"github.com/kglovern/GoSend/CNCServer/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
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
	s.InitializeDB()
}

func (s *CNCServer) InitializeDB() {
	db, err := gorm.Open(sqlite.Open("gosend_rc"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	err = db.AutoMigrate(&models.Macro{})
	if err != nil {
		log.Fatal(err)
	}
	s.DB = db
}

func (s *CNCServer) initializeRoutes() {
	s.Router.HandleFunc("/", s.getHomePage)
}

func (s *CNCServer) Run(port string) {
	log.Fatal(http.ListenAndServe(port, s.Router))
}
