package main

import (
	"database/sql"
	"log"

	"github.com/Kookkla/gofinal/middleware"
	"github.com/Kookkla/gofinal/task"
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

	h := task.Handler{
		DB: db,
	}

	r.GET("/todos", h.GetTodosHandler)
	r.GET("/todos/:id", task.GetTodoByIdHandler)
	r.POST("/todos", task.CreateTodosHandler)
	r.PUT("/todos/:id", task.UpdateTodosHandler)
	r.DELETE("/todos/:id", task.DeleteTodosHandler)

	return r
}

func main() {
	r := setupRouter()
	r.Run(":2009")
}
