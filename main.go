package main

import (
	"github.com/charmbracelet/log"
	"main/routes"
	"main/utils"
	"net/http"
)



func main() {
	utils.IntializeDB()
	mux := http.NewServeMux()
	// API routes (these take precedence over the catch-all due to pattern specificity)
	mux.Handle("/api/", http.StripPrefix("/api", routes.APIRoute()))

	// Catch-all for the SPA: serves from embedded frontend/dist, falling back to index.html
	mux.Handle("/{path...}", SPAHandler("index.html")) // 'dist' is the embedded FS (defined below or imported)
	// Start the server
	log.Info("Server listening on http://localhost:1323")
	if err := http.ListenAndServe(":1323", mux); err != nil {
		log.Info("Server failed:", err.Error())
	}
}
