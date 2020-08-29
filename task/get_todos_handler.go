package task

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetTodosHandler(c *gin.Context) {
	status := c.Query("status")

	stmt, err := h.DB.Prepare("SELECT id, title, status FROM todos")
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	rows, err := stmt.Query()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	todos := []Todo{}
	for rows.Next() {
		t := Todo{}

		err := rows.Scan(&t.ID, &t.Title, &t.Status)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
			return
		}

		todos = append(todos, t)
	}

	tt := []Todo{}

	for _, item := range todos {
		if status != "" {
			if item.Status == status {
				tt = append(tt, item)
			}
		} else {
			tt = append(tt, item)
		}
	}

	c.JSON(http.StatusOK, tt)
}
