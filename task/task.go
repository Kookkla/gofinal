package task

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}
}

type Todo struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Status string `json:"status"`
}

type Handler struct {
	DB *sql.DB
}

func GetTodoByIdHandler(c *gin.Context) {
	id := c.Param("id")

	stmt, err := db.Prepare("SELECT id, name, email,status FROM customers where id=$1")
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	row := stmt.QueryRow(id)

	t := &Todo{}

	err = row.Scan(&t.ID, &t.Name, &t.Email, &t.Status)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, t)
}
