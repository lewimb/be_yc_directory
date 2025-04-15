package startups

import (
	"database/sql"
	"fmt"
	h "github/lewimb/be_yc_directory/handlers"
	"net/http"
)

func RegisterRoutesStartUp(mux *http.ServeMux, db *sql.DB) {

	var handler = h.StartupHandler{
		DB: db,
	}

	// Get All Startup
	mux.HandleFunc("GET /startup", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Successfully get startup data")
	})

	// Get startup by id
	mux.HandleFunc("GET /startup/{id}", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Successfully get startup data")
	})

	// Create startup
	mux.HandleFunc("POST /startup", handler.CreateStartup)

	// Edit startup
	mux.HandleFunc("PUT /startup", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Successfully change startup data")
	})

}
