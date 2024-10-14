package menumemory_backend

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"menumemory-backend/db"
	"net/http"
	"os"
	"time"
)

func main() {
	fmt.Println("Beginning Database Initialization")
	db_, err := sql.Open("sqlite3", "warehouse.db")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	q := db.New(db_)
	fmt.Println("Finished Database Initialization")

	r := gin.Default()

	type PostVisitParams struct {
		Date         string `json:"date" binding:"required"`
		RestaurantId int    `json:"restaurant_id" binding:"required"`
		Time         string `json:"time"`
	}
	r.POST("/visits", func(c *gin.Context) {
		params := PostVisitParams{}
		err := c.BindJSON(&params)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}

		q.CreateVisit(db.CreateVisitParams{
			Date: time.Now().Format("2006-01-02T15:04:05"),
		})
	})

	r.Run() // listen and serve on 0.0.0.0:8080
}
