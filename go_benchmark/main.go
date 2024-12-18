package main

import (
	"fmt"
	"gobenchmark/database"
	"gobenchmark/models"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	envErr := godotenv.Load()
	if envErr != nil {
		e.Logger.Fatal("Error loading .env")
	}

	database.DatabaseConnect()

	e.GET("/story", func(c echo.Context) error {
		isDescAsc := c.QueryParam("sortBy")
		dataGroup := []models.StoryModels{}
		var query string
		
		if isDescAsc == "DESC" {
			query = "SELECT * FROM story ORDER BY story_id DESC LIMIT 20"
		} else {
			query = "SELECT * FROM story ORDER BY story_id ASC LIMIT 20"
		}
		err := database.DBSqlX.Select(&dataGroup, query)
		if err != nil {
			fmt.Println("Query Error:", err)
			return c.JSON(500, map[string]string{"error": "Failed to fetch data"})
		}
		return c.JSON(200, dataGroup)
	})

	e.GET("/store-kategori/:kategori", func(c echo.Context) error {
		var dataGroup []models.StoryModels
		param := c.Param("kategori")
		queryHit := fmt.Sprintf("SELECT * FROM story WHERE story.kategori = %s limit 20", param)
		queryData := database.DBSqlX.Select(dataGroup, queryHit)

		if queryData != nil {
			return c.JSON(200, dataGroup)
		}
		return nil
	})
	e.Logger.Fatal(e.Start(":3000"))
}
