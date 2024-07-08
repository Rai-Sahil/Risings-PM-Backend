package app

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/microsoft"
	"net/http"
	"pm_backend/internals/database"
	"pm_backend/internals/models"
	"strconv"
	"time"
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

var jwtKey = []byte("your-secret-key")

func generateJWT(userID string, name string, email string) (string, error) {
	claims := jwt.MapClaims{
		"userID": userID,
		"name":   name,
		"email":  email,
		"exp":    time.Now().Add(time.Hour * 24).Unix(), // Token expiration time (e.g., 24 hours)
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

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

	user := models.User{
		Name:  userInfo["displayName"].(string),
		Email: userInfo["mail"].(string),
	}

	userID, err := database.InsertUser(user)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to insert user: %v", err), http.StatusInternalServerError)
		return
	}

	// Generate JWT token
	signedToken, err := generateJWT(strconv.FormatInt(userID, 10), user.Name, user.Email)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to generate JWT: %v", err), http.StatusInternalServerError)
		return
	}

	cookie := http.Cookie{
		Name:     "jwt",
		Value:    signedToken,
		Expires:  time.Now().Add(24 * time.Hour), // Adjust as needed
		Path:     "/",
		HttpOnly: true,
		Secure: true,
		SameSite: http.SameSiteNoneMode,
	}
	http.SetCookie(w, &cookie)

	// Set JWT in response header
	response := map[string]interface{}{
		"message": "Authentication successful",
	}
	json.NewEncoder(w).Encode(response)

	// Redirect to frontend URL
	redirectURL := "https://pm-frontend-swart.vercel.app/auth/callback" // Adjust with your frontend URL
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

func getUserDetails(w http.ResponseWriter, r *http.Request) {
	tokenString, err := r.Cookie("jwt")
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	token, err := jwt.Parse(tokenString.Value, func(token *jwt.Token) (interface{}, error) {
		// Make sure that the token method conforms to signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return jwtKey, nil
	})
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	if !token.Valid {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	userID := fmt.Sprintf("%v", claims["userID"])

	var user models.User
	if err := database.DB.First(&user, userID).Error; err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	userDetails := map[string]interface{}{
		"id":    user.ID,
		"name":  user.Name,
		"email": user.Email,
	}

	json.NewEncoder(w).Encode(userDetails)
}

func AuthRoutes(r *mux.Router) {
	r.HandleFunc("/auth/signin", handleMicrosoftLogin).Methods("GET")
	r.HandleFunc("/auth/callback", handleMicrosoftCallback).Methods("GET")
	r.HandleFunc("/auth/user", getUserDetails).Methods("GET")
}
