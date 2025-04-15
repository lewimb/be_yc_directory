package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github/lewimb/be_yc_directory/lib/pkg"
	m "github/lewimb/be_yc_directory/models"
	"log"
	"net/http"
)

type StartupHandler struct {
	DB *sql.DB
}

func (sh *StartupHandler) CreateStartup(w http.ResponseWriter, r *http.Request) {
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

	_, err = sh.DB.Exec("INSERT INTO startup (`title`,`category`,`pitch`,`image`,`slug`,`desc`,`userId` ) VALUES (?,?,?,?,?,?,?)", startup.Title, startup.Category, startup.Pitch, startup.Image, slug, startup.Desc, id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "Succesfully create a startup")
}

func (sh *StartupHandler) GetAllStartup(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Not the correct method", http.StatusMethodNotAllowed)
	}

	var startups []m.Startup
	var profilePics []sql.NullString

	rows, err := sh.DB.Query(`SELECT startup.title,startup.category,startup.pitch,startup.image
	,startup.slug,startup.desc,user.username,user.profile_pic 
	FROM startup INNER JOIN user ON startup.userId=user.id`)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	for rows.Next() {
		var startup m.Startup
		var profilePic sql.NullString
		if err := rows.Scan(&startup.Title, &startup.Category, &startup.Pitch, &startup.Image, &startup.Slug, &startup.Desc, &startup.User.Username, &profilePic); err != nil {
			log.Fatal(err)
		}

		profilePics = append(profilePics, profilePic)
		startups = append(startups, startup)
	}
	rerr := rows.Close()
	if rerr != nil {
		log.Fatal(rerr)
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	for i, pic := range profilePics {
		if pic.Valid {
			startups[i].User.ProfilePic = pic.String
		} else {
			startups[i].User.ProfilePic = ""
		}
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)

	var response = make([]map[string]any, len(startups))

	for i, data := range startups {
		response[i] = map[string]any{
			"title":       data.Title,
			"category":    data.Category,
			"desc":        data.Desc,
			"pitch":       data.Pitch,
			"image":       data.Image,
			"slug":        data.Slug,
			"username":    data.User.Username,
			"profile_pic": data.User.ProfilePic,
		}
	}

	defer json.NewEncoder(w).Encode(response)

}
