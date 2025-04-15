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
	mux.HandleFunc("GET /startup", handler.GetAllStartup)

	// Get startup by slug
	mux.HandleFunc("GET /startup/{slug}", handler.GetStartupBySlug)

	// Create startup
	mux.HandleFunc("POST /startup", handler.CreateStartup)

	// Delete startup
	mux.HandleFunc("DELETE /startup/{slug}", handler.DeleteStartup)

	// Edit startup
	mux.HandleFunc("PUT /startup", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Successfully change startup data")
	})

}
