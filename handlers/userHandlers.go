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

type UserHandler struct {
	DB *sql.DB
}

func (uh *UserHandler) SignUp(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Unable to signup", http.StatusBadRequest)
		panic("Method not POST")
	}

	var user m.SignUpCredential

	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		panic("Unable to decode")
	}

	hashPassword, err := pkg.HashPassword(user.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	_, err = uh.DB.Exec("INSERT INTO user (username,email,password) VALUES (?,?,?)", user.Username, user.Email, hashPassword)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusAccepted)
	fmt.Fprint(w, "Successfully created account")
}

func (uh *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Unsupported http method", http.StatusBadRequest)
		return
	}

	var loginCredential m.LoginCredential
	var user m.User

	err := json.NewDecoder(r.Body).Decode(&loginCredential)

	if err != nil {
		return
	}

	fmt.Println(loginCredential.Email)

	row := uh.DB.QueryRow("SELECT id,username,password,email FROM user WHERE email = ?", loginCredential.Email)

	err = row.Scan(&user.ID, &user.Username, &user.Password, &user.Email)

	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	log.Println("Check part")

	if check := pkg.VerifyPassword(loginCredential.Password, user.Password); !check {
		http.Error(w, "Wrong Password", http.StatusUnauthorized)
		return
	}

	tokenString, err := pkg.CreateToken(user.Email, user.ID)

	if err != nil {
		log.Println(tokenString)
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	defer json.NewEncoder(w).Encode(map[string]any{
		"token": tokenString,
		"user": map[string]interface{}{
			"id":       user.ID,
			"username": user.Username,
			"email":    loginCredential.Email,
		},
	})

}

func (uh *UserHandler) GetUserProfile(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Wrong method", http.StatusMethodNotAllowed)
	}

	var user m.User
	var profilePic sql.NullString

	token := pkg.GetHeader(r.Header.Get("Authorization"))

	if err := pkg.VerifyToken(token); err != nil {
		http.Error(w, "Invalid or Expired Token", http.StatusUnauthorized)
	}

	id := pkg.UnloadToken(token)

	err := uh.DB.QueryRow("SELECT username,email,profile_pic FROM user WHERE id = ?", id).Scan(&user.Username, &user.Email, &profilePic)

	if profilePic.Valid {
		user.ProfilePic = profilePic.String
	} else {
		user.ProfilePic = "" // or nil, or some default string
	}

	switch {
	case err == sql.ErrNoRows:
		w.Header().Set("Content-type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("User not found")

	case err != nil:
		log.Fatalf("query error: %v\n", err)
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)

	defer json.NewEncoder(w).Encode(map[string]any{
		"username":   user.Username,
		"email":      user.Email,
		"ProfilePic": user.ProfilePic,
	})
}

// func add() error {
// 	var err error

// 	row := DB.QueryRow("INSERT INTO ")

// 	return err
// }

// func UserHandler(w http.ResponseWriter, r *http.Request) {
// 	if r.URL.Path != "/user" {
// 		http.Error(w, "404 not found", http.StatusNotFound)
// 	}

// 	switch r.Method {
// 	case "GET":
// 		getUser()
// 	case "POST":
// 		addUser()
// 	case "DELETE":
// 		deleteUser()
// 	case "PUT":
// 		editUser()
// 	}

// }
