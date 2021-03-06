package controllers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/mysql"    //mysql database driver
	_ "github.com/jinzhu/gorm/dialects/postgres" //postgres database driver
	_ "gorm.io/driver/mysql"

	"github.com/fawziridwan/backstack/api/models"
)

type Server struct {
	DB     *gorm.DB
	Router *mux.Router
}

func (server *Server) Initialize(driver, user, password, port, host, name string) {
	var err error

	if driver == "mysql" {
		DBURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", driver, user, password, port, host, name)
		server.DB, err = gorm.Open(driver, DBURL)
		if err != nil {
			fmt.Printf("Cannot connect to %s database", driver)
			log.Fatal("This is the error:", err)
		} else {
			fmt.Printf("We are connected to the %s database", driver)
		}
	}

	if driver == "postgres" {
		DBURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", driver, user, password, port, host, name)
		server.DB, err = gorm.Open(driver, DBURL)
		if err != nil {
			fmt.Printf("Cannot connect to %s database", driver)
			log.Fatal("This is the error:", err)
		} else {
			fmt.Printf("We are connected to the %s database", driver)
		}
	}

	server.DB.Debug().AutoMigrate(&models.User{}, &models.Post{}) //database migration

	server.Router = mux.NewRouter()

	server.initializeRoutes()
}

func (server *Server) run(address string) {
	fmt.Println("Listening to port 8080")
	log.Fatal(http.ListenAndServe(address, server.Router))
}
