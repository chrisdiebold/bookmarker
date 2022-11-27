package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type APIServer struct {
	listenAddr string
	store      Storage
}

func NewAPIServer(listenAddr string, store Storage) *APIServer {
	return &APIServer{
		listenAddr: listenAddr,
		store:      store,
	}
}

func (s *APIServer) Run() {
	router := mux.NewRouter()
	router.HandleFunc("/bookmark", makeHTTPHandleFunc(s.handleBookmark))
	log.Println("Bookmark API server running on port: ", s.listenAddr)
	http.ListenAndServe(s.listenAddr, router)

}

func (s *APIServer) handleBookmark(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "GET" {
		return s.handleListBookmark(w, r)
	}

	if r.Method == "POST" {
		return s.handleCreateBookmark(w, r)
	}

	return nil
}

func (s *APIServer) handleCreateBookmark(w http.ResponseWriter, r *http.Request) error {
	keys, ok := r.URL.Query()["link"]
	if !ok || len(keys[0]) < 1 {
		fmt.Println("URL Param missing")
		panic("missing param")
	}

	fmt.Println(keys[0])
	bookmark := NewBookmark("https://www.youtube.com/watch?v=pwZuNmAzaH8", "GO API Server")
	return WriteJSON(w, http.StatusOK, bookmark)
}

func (s *APIServer) handleListBookmark(w http.ResponseWriter, r *http.Request) error {
	return nil
}

type apiFunc func(http.ResponseWriter, *http.Request) error

func makeHTTPHandleFunc(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			// handle error
			WriteJSON(w, http.StatusBadRequest, ApiError{Error: err.Error()})
		}
	}
}

func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}

type ApiError struct {
	Error string
}
