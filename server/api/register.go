package api

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"

	"go-vue-message-board/models"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func Register(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user models.User

	json.NewDecoder(r.Body).Decode(&user)
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.MinCost)

	if err != nil {
		log.Panic(err)
	}

	user.Password = string(hash)
	db.Create(&user)

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

	w.WriteHeader(http.StatusCreated)
	data := make(map[string]interface{})
	data["user"] = user
	data["token"] = token

	json.NewEncoder(w).Encode(data)
}