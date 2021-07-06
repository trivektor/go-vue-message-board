package api

import (
	"encoding/json"
	"go-vue-message-board/models"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"gorm.io/gorm"
)

func Login(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user models.User 
	json.NewDecoder(r.Body).Decode(&user)

	os.Setenv("ACCESS_SECRET", "jdnfksdmfksd") //this should be in an env file
  atClaims := jwt.MapClaims{}
  atClaims["authorized"] = true	
	atClaims["user_id"] = user.ID
	atClaims["username"] = user.Username
  atClaims["exp"] = time.Now().Add(time.Hour * 8).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
  token, err := at.SignedString([]byte(os.Getenv("ACCESS_SECRET")))

	if err != nil {
		log.Panic(err)
	}

	data := make(map[string]interface{})
	data["user"] = user
	data["token"] = token

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
}