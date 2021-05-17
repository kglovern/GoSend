package CNCServer

import (
	"fmt"
	"net/http"
)

func (s *CNCServer) getHomePage(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "New Model")
}
