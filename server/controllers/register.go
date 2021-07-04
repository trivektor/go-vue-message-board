package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"go-vue-message-board/models"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func RegisterController(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user models.User 
	json.NewDecoder(r.Body).Decode(&user)
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.MinCost)

	if err != nil {
		log.Panic(err)
	}

	user.Password = string(hash)
	db.Create(&user)
}