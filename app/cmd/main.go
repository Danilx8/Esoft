package main

import (
	"esoft/app/bootstrap"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	app := bootstrap.App()

	env := app.Env

	//timeout := time.Duration(env.ContextTimeout) * time.Second

	e := gin.Default()

	//route.Setup(env, timeout, db, gin)
	e.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"ping": "pong",
		})
	})

	e.Run(env.ServerAddress)
}
