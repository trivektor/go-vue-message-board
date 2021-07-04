package main

import (
	"log"
	"net/http"

	"go-vue-message-board/controllers"

	"github.com/gorilla/mux"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type App struct {
	DB *gorm.DB
	Router *mux.Router
}

func (app *App) ConnectToDB() {
	var err error
	app.DB, err = gorm.Open(postgres.Open("host=localhost user=root password=root dbname=go-vue-message-board port=5432 sslmode=disable"))	

	if err != nil {
		log.Panic(err)
	}
}

func (app *App) InitRoutes() {
	app.Router = mux.NewRouter()
	app.Router.HandleFunc("/register", app.handleRequest(controllers.RegisterController)).Methods("POST")
}

func InitApp() {
	app := new(App)
	app.ConnectToDB()
	app.InitRoutes()
	http.ListenAndServe(":8081", app.Router)
}

// https://github.com/mingrammer/go-todo-rest-api-example/blob/master/app/app.go#L89
func (app *App) handleRequest(handler func(db *gorm.DB, w http.ResponseWriter, r *http.Request)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		handler(app.DB, w, r)
	}
}