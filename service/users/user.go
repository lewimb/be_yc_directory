package users

import (
	"database/sql"
	"fmt"
	handler "github/lewimb/be_yc_directory/handlers"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func RegisterRoutesUser(mux *http.ServeMux, db *sql.DB) {

	userHandle := &handler.UserHandler{
		DB: db,
	}

	// Get User Profile
	mux.HandleFunc("GET /users", userHandle.GetUserProfile)
	//SignUp
	mux.HandleFunc("POST /users/signup", userHandle.SignUp)

	// login
	mux.HandleFunc("POST /users/login", userHandle.Login)

	mux.HandleFunc("PUT /users", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Able to change user data")
	})

}
