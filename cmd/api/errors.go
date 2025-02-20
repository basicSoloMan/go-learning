package main

import (
	"log"
	"net/http"
)

func (app *application) internalServerError(w http.ResponseWriter, r *http.Request, err error) {
	log.Printf("internal server error encountered: %s, path: %s, error: %s", r.Method, r.URL.Path, err.Error())

	writeJSONError(w, http.StatusInternalServerError, "Internal Server Error")
}

func (app *application) badRequestResponse(w http.ResponseWriter, r *http.Request, err error) {
	log.Printf("bad request error encountered: %s, path: %s, error: %s", r.Method, r.URL.Path, err.Error())

	writeJSONError(w, http.StatusBadRequest, err.Error())
}

func (app *application) notFoundResponse(w http.ResponseWriter, r *http.Request, err error) {
	log.Printf("resource not found: %s, path: %s, error: %s", r.Method, r.URL.Path, err.Error())

	writeJSONError(w, http.StatusNotFound, "resource not found")
}
