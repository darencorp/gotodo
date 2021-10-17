package utils

import (
	"net/http"
)

func writeNotAllowed(w http.ResponseWriter) {
	w.WriteHeader(http.StatusMethodNotAllowed)
	w.Write([]byte("Method not allowed"))
}

func Get(wrapped func(http.ResponseWriter, *http.Request)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			wrapped(w, r)
		} else {
			writeNotAllowed(w)
		}
	}
}

func Post(wrapped func(http.ResponseWriter, *http.Request)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			wrapped(w, r)
		} else {
			writeNotAllowed(w)
		}
	}
}

func Patch(wrapped func(http.ResponseWriter, *http.Request)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPatch {
			wrapped(w, r)
		} else {
			writeNotAllowed(w)
		}
	}
}

func Delete(wrapped func(http.ResponseWriter, *http.Request)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodDelete {
			wrapped(w, r)
		} else {
			writeNotAllowed(w)
		}
	}
}
