package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"pm_backend/internals/database"
	"pm_backend/internals/models"
	"strconv"
	"time"
)

var jwtKey = []byte("risings-pm")

type CustomClaims struct {
	Name string `json:"name"`
	Mail string `json:"mail"`
	jwt.StandardClaims
}

func AddUserHandler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	var newUser models.User
	err = json.Unmarshal(body, &newUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	id, err := database.InsertUser(newUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	newUser.ID = id // Ensure the ID is set in the newUser object

	token, err := generateJWT(newUser.Name, newUser.Email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:    "accessToken",
		Value:   token,
		Expires: time.Now().Add(24 * time.Hour),
		Path:    "/",
	})

	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(struct {
		User  models.User `json:"user"`
		Token string      `json:"token"`
	}{
		User:  newUser,
		Token: token,
	}); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func generateJWT(name string, mail string) (string, error) {
	claims := CustomClaims{
		Name: name,
		Mail: mail,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    "your-issuer",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(jwtKey)
	if err != nil {
		return "", fmt.Errorf("failed to generate JWT: %v", err)
	}

	return signedToken, nil
}

func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	users, err := database.GetAllUsers()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = json.NewEncoder(w).Encode(users)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func SetUserStatusHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	userIdStr := mux.Vars(r)["user_id"]
	userId, err := strconv.ParseInt(userIdStr, 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(`{"message": "invalid user id"}`))
	}

	status := mux.Vars(r)["status"]
	if status == "" {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(`{"message": "status is required"}`))
		return
	}

	err = database.SetUserStatus(userId, status)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(`{"message": "invalid status value"}`))
	}

	w.WriteHeader(http.StatusOK)

}
