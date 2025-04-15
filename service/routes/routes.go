package routes

import (
	"database/sql"
	"github/lewimb/be_yc_directory/service/startups"
	"github/lewimb/be_yc_directory/service/users"
	"net/http"
)

type Handler struct{}

func (handler *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {}

func RegisteredRoutes(db *sql.DB) *http.ServeMux {
	mux := http.NewServeMux()
	startups.RegisterRoutesStartUp(mux, db)
	users.RegisterRoutesUser(mux, db)

	return mux
}
