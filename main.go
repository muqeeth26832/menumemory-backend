package main

import (
	"database/sql"
	"fmt"
	"menumemory-backend/db"
	"net/http"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

type CreateDishRequest struct {
	Name string `json:"name" binding:"required"`
}

func SetupApp() *gin.Engine {
	fmt.Println("Beginning Database Initialization")
	db_, err := sql.Open("sqlite3", "warehouse.db")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	q := db.New(db_)
	fmt.Println("Finished Database Initialization")

	r := gin.Default()

	//Allow all cors origins
	r.Use(cors.Default())

	r.StaticFile("/openapi.json", "./openapi.json")
	r.StaticFile("/openapi.yaml", "./openapi.yaml")

	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "Pong Uwu")
	})

	r.GET("/restaurants", func(c *gin.Context) {
		search_term := c.Query("search_term")
		if search_term == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "search_term is required"})
		}

		restaurants, err := q.GetRestaurantsLike(c, "%"+search_term+"%")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}

		c.JSON(http.StatusOK, gin.H{
			"restaurants": restaurants,
		})
	})

	r.POST("/dishes", func(c *gin.Context) {
		var req CreateDishRequest

		if err := c.ShouldBindBodyWithJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Name is required"})
			return
		}

		dish, err := q.CreateDish(c, req.Name)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}

		c.JSON(http.StatusOK, gin.H{
			"dish": dish,
		})
	})

	return r
}

func main() {
	r := SetupApp()

	r.Run() // listen and serve on 0.0.0.0:8080
}
