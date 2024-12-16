package main

import (
	"esoft/app/bootstrap"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	app, err := bootstrap.App()
	if err != nil {
		fmt.Printf("Get error while start app: %v", err)
		return
	}

	//timeout := time.Duration(env.ContextTimeout) * time.Second

	e := gin.Default()

	//route.Setup(env, timeout, db, gin)
	e.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"ping": "pong",
		})
	})

	e.Run(fmt.Sprintf(":%s", app.Config.Server.Port))
}
