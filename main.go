package main

import (
	"database/sql"
	"log"

	"github.com/Kookkla/gofinal/customerservice"
	"github.com/Kookkla/gofinal/middleware"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("postgres", "postgres://vttkxspt:sjA5CdRG1tepOQye8KB1ZMsPjQZ273V9@lallah.db.elephantsql.com:5432/vttkxspt")
	if err != nil {
		log.Fatal(err)
	}
}

func setupRouter() *gin.Engine {
	r := gin.Default()

	r.Use(middleware.Auth)

	h := customerservice.Handler{
		DB: db,
	}

	r.GET("/customers", h.GetCustomersHandler)
	r.GET("/customers/:id", customerservice.GetCustomersByIdHandler)

	r.POST("/customers", customerservice.CreateCustomersHandler)
	r.PUT("/customers/:id", customerservice.UpdateCustomersHandler)
	r.DELETE("/customers/:id", customerservice.DeleteCustomersHandler)

	return r
}

func main() {
	r := setupRouter()
	r.Run(":2009")
}
