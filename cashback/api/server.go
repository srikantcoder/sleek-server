package api

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"sleep.com/cashback/clients"
	"sleep.com/cashback/entities"
)

type Server struct {
	*mux.Router

	deals []entities.Deal

	fileManager *clients.FileManager

	mySQLManager *clients.MySQLManager
}

func NewServer() (*Server, error) {
	mySQLManager, err := clients.NewMySQLManager()
	if err != nil {
		return nil, err
	}

	s := &Server{
		Router:       mux.NewRouter(),
		deals:        []entities.Deal{},
		fileManager:  clients.NewFileManager(),
		mySQLManager: mySQLManager,
	}
	s.routes()
	return s, nil
}

func (s *Server) routes() {
	s.HandleFunc("/activate-deal/{id}", s.activateDeal()).Methods("POST")
	s.HandleFunc("/deals", s.getDeals()).Methods("GET")
	s.HandleFunc("/deals-status", s.getDealsStatus()).Methods("GET")
}

func (s *Server) getDeals() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		deals, err := s.fileManager.GetDeals()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		err = json.NewEncoder(w).Encode(deals)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func (s *Server) activateDeal() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr, _ := mux.Vars(r)["id"]
		dealID, err := uuid.Parse(idStr)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		deals, err := s.fileManager.GetDeals()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		dealFound := false
		for _, deal := range deals {
			if deal.ID == dealID {
				dealFound = true
			}
		}

		if !dealFound {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		s.mySQLManager.ActivateDeal(dealID)
		w.WriteHeader(http.StatusOK)
	}
}

func (s *Server) getDealsStatus() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		dealsStatus := s.mySQLManager.GetDealsStatus()

		w.Header().Set("Content-Type", "application/json")
		err := json.NewEncoder(w).Encode(dealsStatus)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}
