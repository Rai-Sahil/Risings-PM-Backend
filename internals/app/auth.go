package app

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/microsoft"
	"net/http"
	"pm_backend/internals/database"
	"pm_backend/internals/models"
)

var (
	oauth2Config = &oauth2.Config{
		ClientID:     "9e5a29ed-0e39-4234-86e8-0a7f9deac50e",
		ClientSecret: "b_T8Q~iR~jGtJKVkoXoacEBqZeBbXX_.2UktBa1y",
		RedirectURL:  "http://localhost:8080/auth/callback",
		Scopes: []string{
			"https://graph.microsoft.com/User.Read",
		},
		Endpoint: microsoft.AzureADEndpoint("common"),
	}
	oauth2StateString = "random"
)

func handleMicrosoftLogin(w http.ResponseWriter, r *http.Request) {
	url := oauth2Config.AuthCodeURL(oauth2StateString)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func handleMicrosoftCallback(w http.ResponseWriter, r *http.Request) {
	if r.URL.Query().Get("state") != oauth2StateString {
		http.Error(w, "State is not valid", http.StatusBadRequest)
		return
	}

	code := r.URL.Query().Get("code")
	token, err := oauth2Config.Exchange(context.Background(), code)
	if err != nil {
		http.Error(w, fmt.Sprintf("Code exchange failed: %v", err), http.StatusInternalServerError)
		return
	}

	userInfo, err := getUserInfo(token.AccessToken)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed getting user info: %v", err), http.StatusInternalServerError)
		return
	}

	// Create a new user object with the retrieved information
	user := models.User{
		Name:  userInfo["displayName"].(string),
		Email: userInfo["mail"].(string),
	}

	// Insert the user into the database
	userID, err := database.InsertUser(user)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to insert user: %v", err), http.StatusInternalServerError)
		return
	}

	// Redirect to the callback URL with the user details and userID
	redirectURL := fmt.Sprintf("http://localhost:5173/auth/callback?name=%s&email=%s&userID=%d", user.Name, user.Email, userID)
	http.Redirect(w, r, redirectURL, http.StatusTemporaryRedirect)
}

func getUserInfo(accessToken string) (map[string]interface{}, error) {
	url := "https://graph.microsoft.com/v1.0/me"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create HTTP request: %v", err)
	}

	req.Header.Set("Authorization", "Bearer "+accessToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to make request to Microsoft Graph API: %v", err)
	}
	defer resp.Body.Close()

	var userInfo map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&userInfo); err != nil {
		return nil, fmt.Errorf("failed to decode user info response: %v", err)
	}

	return userInfo, nil
}

func AuthRoutes(r *mux.Router) {
	r.HandleFunc("/auth/signin", handleMicrosoftLogin).Methods("GET")
	r.HandleFunc("/auth/callback", handleMicrosoftCallback).Methods("GET")
}
