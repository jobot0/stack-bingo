package main

import (
	"net/http"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg/v9"
	"stack-bingo/core"
	_dbDumbRepository "stack-bingo/data"
)

type jsonDumb struct {
	Name string `json:"name" binding:"required"`
}

func main() {
	db := pg.Connect(&pg.Options{
		Database: "stack-bingo",
		User: "postgres",
		Password: "Pass2020!",
	})

	dumbRepo := _dbDumbRepository.NewDbDumbRepository(db)
	dumbu := core.NewCreateDumbUsecase(dumbRepo)
	// Set the router as the default one shipped with Gin
	router := gin.Default()

	// Serve frontend static files
	router.Use(static.Serve("/", static.LocalFile("./front/build", true)))

	// Setup route group for the API
	api := router.Group("/api")
	{
		api.POST("/dumb", func(c *gin.Context) {
			var toto jsonDumb
			c.BindJSON(&toto)
			d := core.Dumb{Name: toto.Name}
			dumbu.Execute(&d)
		})
		api.GET("/", func(c *gin.Context){
			c.JSON(http.StatusOK, gin.H{
				"message": "pong",
			})
		})
	}

	//defer db.Close()
	// Start and run the server
	router.Run(":5000")
}

