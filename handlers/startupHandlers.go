package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github/lewimb/be_yc_directory/lib/pkg"
	m "github/lewimb/be_yc_directory/models"
	"net/http"
)

type StartupHandler struct {
	DB *sql.DB
}

func (handler *StartupHandler) CreateStartup(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Not a post method", http.StatusMethodNotAllowed)
	}
	var startup m.Startup

	err := json.NewDecoder(r.Body).Decode(&startup)

	token := pkg.GetHeader(r.Header.Get("Authorization"))
	id := pkg.UnloadToken(token)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	slug := pkg.CreateSlug(startup.Title)

	_, err = handler.DB.Exec("INSERT INTO startup (`title`,`category`,`pitch`,`image`,`slug`,`desc`,`userId` ) VALUES (?,?,?,?,?,?,?)", startup.Title, startup.Category, startup.Pitch, startup.Image, slug, startup.Desc, id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "Succesfully create a startup")
}
