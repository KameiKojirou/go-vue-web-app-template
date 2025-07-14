package routes

import (
	"net/http"
)

func APIRoute() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/world", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("World"))
	})
	mux.HandleFunc("/login", Login)
	mux.HandleFunc("/logout", Logout)
	// catch all that will return 404 on any api route that doesn't exist
	mux.HandleFunc("/{path...}", func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Not Found", http.StatusNotFound)
	})
	return mux
}